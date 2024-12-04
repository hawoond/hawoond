package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hawoond/hawoond/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

type Jet struct {
	Header    Header
	Claims    Claims
	Blacklist DBBlacklist
}

// Header 구조체 정의
type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

// Claims 구조체 정의
type Claims struct {
	Issuer    string                 `json:"iss,omitempty"`
	Subject   string                 `json:"sub,omitempty"`
	Audience  string                 `json:"aud,omitempty"`
	ExpiresAt int64                  `json:"exp,omitempty"`
	NotBefore int64                  `json:"nbf,omitempty"`
	IssuedAt  int64                  `json:"iat,omitempty"`
	ID        string                 `json:"jti,omitempty"`
	TokenType string                 `json:"typ,omitempty"` // 토큰 유형 추가
	Data      map[string]interface{} `json:"data,omitempty"`
}

// SigningMethod 인터페이스 정의
type SigningMethod interface {
	Sign(message string) (string, error)
	Verify(message, signature string) error
	Alg() string
}

// HS256 알고리즘 구조체
type HS256 struct {
	Secret []byte
}

func (h *HS256) Sign(message string) (string, error) {
	mac := hmac.New(sha256.New, h.Secret)
	mac.Write([]byte(message))
	signature := mac.Sum(nil)
	return utils.Base64UrlEncode(signature), nil
}

func (h *HS256) Verify(message, signature string) error {
	expectedSig, err := h.Sign(message)
	if err != nil {
		return err
	}
	if !hmac.Equal([]byte(signature), []byte(expectedSig)) {
		return errors.New("서명이 유효하지 않습니다")
	}
	return nil
}

func (h *HS256) Alg() string {
	return "HS256"
}

// 에러 타입 정의
var (
	ErrInvalidToken     = errors.New("유효하지 않은 토큰입니다")
	ErrExpiredToken     = errors.New("토큰이 만료되었습니다")
	ErrNotValidYetToken = errors.New("토큰 사용이 허가되지 않았습니다")
	ErrInvalidSignature = errors.New("서명이 유효하지 않습니다")
	ErrInvalidAlgorithm = errors.New("지원하지 않는 알고리즘입니다")
	ErrInvalidKey       = errors.New("유효하지 않은 키입니다")
	ErrClaimsValidation = errors.New("클레임 검증에 실패했습니다")
)

// 토큰 생성 함수
func (j *Jet) CreateToken(claims Claims, method SigningMethod) (string, error) {
	header := Header{
		Alg: method.Alg(),
		Typ: "JWT",
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	headerEncoded := utils.Base64UrlEncode(headerJSON)
	claimsEncoded := utils.Base64UrlEncode(claimsJSON)

	unsignedToken := headerEncoded + "." + claimsEncoded
	signature, err := method.Sign(unsignedToken)
	if err != nil {
		return "", err
	}

	token := unsignedToken + "." + signature
	return token, nil
}

// VerifyToken 함수 수정
func (j *Jet) VerifyToken(token string, method SigningMethod, options *VerificationOptions, blacklist Blacklist) (*Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, ErrInvalidToken
	}

	headerEncoded, claimsEncoded, signature := parts[0], parts[1], parts[2]

	unsignedToken := headerEncoded + "." + claimsEncoded

	// 서명 검증
	if err := method.Verify(unsignedToken, signature); err != nil {
		return nil, ErrInvalidSignature
	}

	// 헤더 파싱
	headerJSON, err := utils.Base64UrlDecode(headerEncoded)
	if err != nil {
		return nil, err
	}

	var header Header
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, err
	}

	// 알고리즘 검증
	if header.Alg != method.Alg() {
		return nil, ErrInvalidAlgorithm
	}

	// 클레임 파싱
	claimsJSON, err := utils.Base64UrlDecode(claimsEncoded)
	if err != nil {
		return nil, err
	}

	var claims Claims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, err
	}

	// 블랙리스트 확인
	if blacklist != nil {
		exists, err := blacklist.Contains(claims.ID)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("블랙리스트에 등록된 토큰입니다")
		}
	}

	// 클레임 검증
	now := time.Now().Unix()

	// 유예 기간 적용
	gracePeriod := int64(0)
	if options != nil {
		gracePeriod = options.GracePeriod
	}

	if claims.ExpiresAt != 0 && now > claims.ExpiresAt+gracePeriod {
		return nil, ErrExpiredToken
	}

	if claims.NotBefore != 0 && now < claims.NotBefore-gracePeriod {
		return nil, ErrNotValidYetToken
	}

	if claims.IssuedAt != 0 && now < claims.IssuedAt-gracePeriod {
		return nil, ErrClaimsValidation
	}

	// Issuer 검증
	if options != nil && len(options.ValidIssuers) > 0 {
		validIssuer := false
		for _, issuer := range options.ValidIssuers {
			if claims.Issuer == issuer {
				validIssuer = true
				break
			}
		}
		if !validIssuer {
			return nil, errors.New("유효하지 않은 Issuer입니다")
		}
	}

	// Audience 검증
	if options != nil && len(options.ValidAudiences) > 0 {
		validAudience := false
		for _, audience := range options.ValidAudiences {
			if claims.Audience == audience {
				validAudience = true
				break
			}
		}
		if !validAudience {
			return nil, errors.New("유효하지 않은 Audience입니다")
		}
	}

	return &claims, nil
}

// JTI 생성 함수
func (j *Jet) GenerateJTI() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// VerificationOptions 구조체 정의
type VerificationOptions struct {
	ValidIssuers   []string
	ValidAudiences []string
	GracePeriod    int64 // 초 단위
}

// Blacklist 인터페이스
type Blacklist interface {
	Add(jti string, expiresAt int64) error
	Remove(jti string) error
	Contains(jti string) (bool, error)
}

// DBBlacklist 구조체
type DBBlacklist struct {
	db *sql.DB
}

// DBBlacklist 생성 함수
func (j *Jet) NewDBBlacklist(dataSourceName string) (*DBBlacklist, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	// 테이블 생성
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS blacklist (
        jti TEXT PRIMARY KEY,
        expired_at INTEGER
    );
    `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, err
	}

	return &DBBlacklist{db: db}, nil
}

// 블랙리스트에 추가
func (b *DBBlacklist) Add(jti string, expiresAt int64) error {
	query := `INSERT OR REPLACE INTO blacklist (jti, expired_at) VALUES (?, ?)`
	_, err := b.db.Exec(query, jti, expiresAt)
	return err
}

// 블랙리스트에서 삭제
func (b *DBBlacklist) Remove(jti string) error {
	query := `DELETE FROM blacklist WHERE jti = ?`
	_, err := b.db.Exec(query, jti)
	return err
}

// 블랙리스트에 존재하는지 확인
func (b *DBBlacklist) Contains(jti string) (bool, error) {
	query := `SELECT expired_at FROM blacklist WHERE jti = ?`
	row := b.db.QueryRow(query, jti)

	var expiredAt int64
	err := row.Scan(&expiredAt)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}

	// 만료된 블랙리스트 항목은 자동으로 삭제
	if time.Now().Unix() > expiredAt {
		_ = b.Remove(jti)
		return false, nil
	}

	return true, nil
}

// 블랙리스트 만료된 항목 삭제 (예시로 사용)
func (b *DBBlacklist) CleanUp() error {
	query := `DELETE FROM blacklist WHERE expired_at <= ?`
	_, err := b.db.Exec(query, time.Now().Unix())
	return err
}
