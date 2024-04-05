package main

import "fmt"

type SocialNetwork struct {
	g *Graph
}

func NewSocialNetwork(g *Graph) *SocialNetwork {
	return &SocialNetwork{g: g}
}

func (sn *SocialNetwork) AddUser(val int) {
	sn.g.Vertices[val] = &Vertex{Val: val, Edges: map[int]*Edge{}}
}

func (sn *SocialNetwork) MakeFriends(val1, val2 int) {
	sn.g.Vertices[val1].Edges[val2] = &Edge{Weight: 1, Vertex: sn.g.Vertices[val2]}
	sn.g.Vertices[val2].Edges[val1] = &Edge{Weight: 1, Vertex: sn.g.Vertices[val1]}
}

func (sn *SocialNetwork) DisplayNetwork() {
	for _, user := range sn.g.Vertices {
		fmt.Printf("\nuser:%d\nFriends:", user.Val)
		for friend, _ := range sn.g.Vertices[user.Val].Edges {
			fmt.Printf(" %d,", friend)
		}
	}
}

func (sn *SocialNetwork) GetMostFollowedUser() {}
