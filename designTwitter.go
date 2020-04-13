package leetcode

import "container/heap"

var timestamp int

type Twitter struct {
	Users map[int]*User
}

type Tweet struct {
	ID   int
	Time int
	Next *Tweet
}

func NewTweet(id int, time int) *Tweet {
	return &Tweet{ID: id, Time: time}
}

type PriorityQueue []*Tweet

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Time > pq[j].Time
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	item := old[len(old)-1]
	*pq = old[:len(old)-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Tweet)
	*pq = append(*pq, item)
}

type User struct {
	ID       int
	Followed map[int]bool
	Head     *Tweet
}

func NewUser(id int) *User {
	user := &User{}
	user.ID = id
	user.Followed = make(map[int]bool)
	user.follow(id)
	return user
}

func (u *User) follow(id int) {
	u.Followed[id] = true
}

func (u *User) unFollow(id int) {
	if id != u.ID && u.Followed[id] {
		delete(u.Followed, id)
	}
}

func (u *User) post(id int) {
	tweet := NewTweet(id, timestamp)
	timestamp++
	tweet.Next = u.Head
	u.Head = tweet
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{Users: make(map[int]*User)}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.Users[userId]; !ok {
		this.Users[userId] = NewUser(userId)
	}
	this.Users[userId].post(tweetId)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	if _, ok := this.Users[userId]; !ok {
		return nil
	}
	user := this.Users[userId]
	pq := make(PriorityQueue, 0)
	for k, _ := range user.Followed {
		twt := this.Users[k].Head
		if twt != nil {
			pq = append(pq, twt)
		}
	}
	heap.Init(&pq)
	var ans []int
	for pq.Len() != 0 {
		if len(ans) == 10 {
			break
		}
		twt := heap.Pop(&pq).(*Tweet)
		ans = append(ans, twt.ID)
		if twt.Next != nil {
			heap.Push(&pq, twt.Next)
		}
	}
	return ans
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.Users[followerId]; !ok {
		this.Users[followerId] = NewUser(followerId)
	}
	if _, ok := this.Users[followeeId]; !ok {
		this.Users[followeeId] = NewUser(followeeId)
	}
	this.Users[followerId].follow(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if user, ok := this.Users[followerId]; ok {
		user.unFollow(followeeId)
	}
}
