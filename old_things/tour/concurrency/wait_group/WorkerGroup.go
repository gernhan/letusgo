package wait_group

import "sync"

type WorkerGroup struct {
	commitment *sync.WaitGroup
}

func (w *WorkerGroup) Work(group ...*Worker) {
	var commitment sync.WaitGroup
	w.commitment = &commitment
	for _, worker := range group {
		worker.WithCommitment(&commitment)
		go worker.Work()
	}
	commitment.Wait()
}
