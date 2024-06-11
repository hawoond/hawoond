package worker

import (
	"fmt"
	"sync"
	"time"
)

// 기본 우선순위를 설정
const defaultPriority = 0

// Job 구조체 정의: 작업을 함수로 정의
type Job struct {
	priority int          // 작업 우선순위
	execute  func() error // 실행할 작업
}

// PriorityQueue 구조체 정의
type PriorityQueue struct {
	jobs []*Job
	lock sync.Mutex
}

// PriorityQueue 초기화 함수
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		jobs: []*Job{},
	}
}

// 작업을 힙에 추가하는 함수
func (pq *PriorityQueue) Push(job *Job) {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if job.priority == 0 {
		job.priority = defaultPriority
	}
	pq.jobs = append(pq.jobs, job)
	pq.up(len(pq.jobs) - 1)
}

// 힙에서 작업을 꺼내는 함수
func (pq *PriorityQueue) Pop() *Job {
	pq.lock.Lock()
	defer pq.lock.Unlock()
	if len(pq.jobs) == 0 {
		return nil
	}
	pq.swap(0, len(pq.jobs)-1)
	job := pq.jobs[len(pq.jobs)-1]
	pq.jobs = pq.jobs[:len(pq.jobs)-1]
	pq.down(0)
	return job
}

// 힙의 상향 조정 함수
func (pq *PriorityQueue) up(index int) {
	for {
		parent := (index - 1) / 2
		if parent == index || pq.jobs[parent].priority <= pq.jobs[index].priority {
			break
		}
		pq.swap(parent, index)
		index = parent
	}
}

// 힙의 하향 조정 함수
func (pq *PriorityQueue) down(index int) {
	for {
		child := 2*index + 1
		if child >= len(pq.jobs) || child < 0 {
			break
		}
		if right := child + 1; right < len(pq.jobs) && pq.jobs[right].priority < pq.jobs[child].priority {
			child = right
		}
		if pq.jobs[child].priority >= pq.jobs[index].priority {
			break
		}
		pq.swap(child, index)
		index = child
	}
}

// 힙의 두 요소를 교환하는 함수
func (pq *PriorityQueue) swap(i, j int) {
	pq.jobs[i], pq.jobs[j] = pq.jobs[j], pq.jobs[i]
}

// Worker 구조체 정의
type Worker struct {
	pq      *PriorityQueue
	quit    chan struct{}
	wg      *sync.WaitGroup
	running bool
}

// Worker 초기화 함수
func NewWorker(pq *PriorityQueue) *Worker {
	return &Worker{
		pq:      pq,
		quit:    make(chan struct{}),
		wg:      &sync.WaitGroup{},
		running: false,
	}
}

// Worker 시작 함수
func (w *Worker) Start(wt ...int) {
	if len(wt) == 0 {
		wt[0] = 100
	}
	wait := wt[0]
	w.running = true
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			select {
			case <-w.quit:
				return
			default:
				job := w.pq.Pop()
				if job != nil {
					// 작업 처리
					err := job.execute()
					if err != nil {
						fmt.Printf("Error executing job: %v\n", err)
					} else {
						fmt.Printf("Successfully executed job with priority %d\n", job.priority)
					}
				}
				// 잠시 대기
				time.Sleep(time.Duration(wait) * time.Millisecond)
			}
		}
	}()
}

// Worker 종료 함수
func (w *Worker) Stop() {
	close(w.quit)
	w.wg.Wait()
	w.running = false
}
