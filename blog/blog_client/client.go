package main

import (
	"context"
	"fmt"
	"log"

	"github.com/xvbnm48/go-grpc-learning/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")
	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create blog
	fmt.Println("Created BlogServiceClient")
	blog := &blogpb.Blog{
		AuthorId: "Sakura",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	createBlogs, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog was created: %v\n", createBlogs)
	blogID := createBlogs.GetBlog().GetId()
	// READ BLOG
	fmt.Println("Reading Blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5f18f8f8c9f8f8f8f8f8f8f8"})
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogResErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogResErr != nil {
		fmt.Printf("Error happened while reading: %v\n", readBlogResErr)
	}
	fmt.Printf("Blog was read: %v\n", readBlogRes)
	// update blog
	newBlog := &blogpb.Blog{Id: blogID, AuthorId: "takuya", Title: "i can change this title blog", Content: "the this content is changed"}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if updateErr != nil {
		fmt.Printf("Error hapened while updating blog: %v", updateErr)
	}
	fmt.Printf("Blog was read : %v\n", updateRes)

	// delete blog
	fmt.Println("Deleting Blog")
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if deleteRes != nil {
		fmt.Printf("Blog was deleted: %v\n", deleteErr)
	}

	fmt.Printf("Blog was deleted: %v\n", deleteRes)
}
