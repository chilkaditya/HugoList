package services

import (
    "context"
    "fmt"
    "youtube-playlist-automation/utils"

    "google.golang.org/api/youtube/v3"
)

func CreatePlaylistAndAddSongs(title, description string, songs []string) error {
    service := utils.GetClient()
    ctx := context.Background()

    // 1. Create Playlist
    playlist := &youtube.Playlist{
        Snippet: &youtube.PlaylistSnippet{
            Title:       title,
            Description: description,
        },
        Status: &youtube.PlaylistStatus{
            PrivacyStatus: "private",
        },
    }

    createCall := service.Playlists.Insert([]string{"snippet,status"}, playlist)
    createdPlaylist, err := createCall.Context(ctx).Do()
    if err != nil {
        return fmt.Errorf("error creating playlist: %v", err)
    }

    // 2. Search each song and add to playlist
    for _, song := range songs {
        // Search video
        searchCall := service.Search.List([]string{"id,snippet"}).Q(song).MaxResults(1)
        searchResult, err := searchCall.Do()
        if err != nil || len(searchResult.Items) == 0 {
            continue
        }
        videoId := searchResult.Items[0].Id.VideoId

        // Add video to playlist
        playlistItem := &youtube.PlaylistItem{
            Snippet: &youtube.PlaylistItemSnippet{
                PlaylistId: createdPlaylist.Id,
                ResourceId: &youtube.ResourceId{
                    Kind:    "youtube#video",
                    VideoId: videoId,
                },
            },
        }

        _, err = service.PlaylistItems.Insert([]string{"snippet"}, playlistItem).Context(ctx).Do()
        if err != nil {
            fmt.Println("Error adding video:", song)
            continue
        }
    }

    return nil
}
