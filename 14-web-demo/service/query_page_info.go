package service

import (
	"LearnGoProject/14-web-demo/repository"
	"errors"
	"sync"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}
type QueryInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo
	topic    *repository.Topic
	posts    []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryInfoFlow(topicId).Do()
}
func NewQueryInfoFlow(topId int64) *QueryInfoFlow {
	return &QueryInfoFlow{
		topicId: topId,
	}
}
func (f *QueryInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}
func (f *QueryInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil

}
func (f *QueryInfoFlow) prepareInfo() error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryPostsByParentId(f.topicId)
		f.posts = posts
	}()
	wg.Wait()
	return nil

}
func (f *QueryInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil

}
