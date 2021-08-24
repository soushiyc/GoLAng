package main

import "fmt"

type tvseries struct{
	Name string
	Service string
	next *tvseries
}

type playlist struct{
	Name		string
	head 		*tvseries
	nowWatching *tvseries
}

func createPlaylist(name string) *playlist{
	return &playlist{
		Name: name,
	}
}

func (p *playlist) addSeries(name, service string) error{
	show := &tvseries{
		Name: name,
		Service: service,
	}
	if p.head == nil{
		p.head = show
	} else {
		currentNode := p.head
		for currentNode.next !=nil{
			currentNode = currentNode.next
		}
		currentNode.next = show
	}
	return nil
}

func (p *playlist) showAllSeries() error{
	currentNode := p.head
	if currentNode == nil{
		fmt.Println("The current playlist is empty")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	return nil
}

func (p *playlist) startPlaying() *tvseries{
	p.nowWatching = p.head
	return p.nowWatching
}

func (p *playlist) nextSeries() *tvseries{
	p.nowWatching = p.nowWatching.next
	return p.nowWatching
}

func main(){
	playlistName := "My Series"
	mySeries := createPlaylist(playlistName)
	fmt.Println("Playlist created")
	fmt.Println()

	var i bool = true
	var name, service string
	fmt.Println("Please enter the shows and its service you would like to add to the playlist.\nEnter 0 and 0 whe you are done")
	for i == true {
		fmt.Scan(&name)
		fmt.Scan(&service)
		if name == "0" && service =="0" {
			i = false
		} else {
			mySeries.addSeries(name, service)
		}
	}


	mySeries.startPlaying()
	fmt.Printf("Now watching: %s, %s", mySeries.nowWatching.Name, mySeries.nowWatching.Service)
	fmt.Println()

 	for mySeries.nowWatching.next != nil{
		mySeries.nextSeries()
 		fmt.Println("Changing to nest series: ")
		fmt.Printf("Now watching: %s, %s", mySeries.nowWatching.Name, mySeries.nowWatching.Service)
 		fmt.Println()
	}
}

/*

*/

/*
	fmt.Print("Adding songs to the playlist...\n")
	mySeries.addSeries("Ophelia", "The Lumineers")
	mySeries.addSeries("Shape of you", "Ed Sheeran")
	mySeries.addSeries("Stubborn Love", "The Lumineers")
	mySeries.addSeries("Feels", "Calvin Harris")
	mySeries.addSeries("You should be dancing", "Beco")
	fmt.Println("Displaying all shows in the playlist")
	mySeries.showAllSeries()
	fmt.Println()
*/