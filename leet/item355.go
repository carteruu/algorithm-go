package leet

type Twitter struct {
	follow []map[int]struct{}
	tweet  [][2]int
}

func ConstructorTwitter() Twitter {
	return Twitter{
		follow: make([]map[int]struct{}, 501),
		tweet:  make([][2]int, 10001),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweet = append(this.tweet, [2]int{userId, tweetId})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	feed := make([]int, 0, 10)
	for i := len(this.tweet) - 1; i >= 0; i-- {
		followeeId := this.tweet[i][0]
		if _, ok := this.follow[userId][followeeId]; followeeId == userId || ok {
			feed = append(feed, this.tweet[i][1])
			if len(feed) == 10 {
				break
			}
		}
	}
	return feed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.follow[followerId] == nil {
		this.follow[followerId] = make(map[int]struct{})
	}
	this.follow[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	delete(this.follow[followerId], followeeId)
}
