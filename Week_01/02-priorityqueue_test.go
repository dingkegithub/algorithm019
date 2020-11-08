package Week_01

import "testing"

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()

	pq.Add(1)
	pq.Add(2)
	pq.Add(3)

	el := pq.Peek()
	if el == nil {
		t.Log("peek element err")
		t.FailNow()
	}

	if el.(int) != 3 {
		t.Log("peek failed, want 3, but: ", el.(int))
		t.FailNow()
	}

	el = pq.Poll()
	if el == nil {
		t.Log("poll element err")
		t.FailNow()
	}

	if el.(int) != 3 {
		t.Log("poll failed, want 3, but: ", el.(int))
		t.FailNow()
	}

	pq.Add(6)

	el, err := pq.Remove()
	if err != nil {
		t.Error("remove element failed: ", err)
		t.FailNow()
	}

	if el.(int) != 6 {
		t.Log("poll failed, want 6, but: ", el.(int))
		t.FailNow()
	}

}
