package queries

import (
	"test/dbConnection"
	"test/models"
)


// GetAlbumsQuery returns a list of all albums in the database.
func GetAlbumsQuery() ([]models.Album, error) {
	// Execute the SELECT statement
	rows, err := dbConnection.Connection().Query("SELECT id, title, artist, price FROM album")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Iterate through the result set
	albums := []models.Album{}
	for rows.Next() {
		var a models.Album
		err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, a)
	}
	// Return the list of albums
	return albums, nil
}



// PostAlbumQuery add an array of album to the database
func PostAlbumQuery(a []models.Album) ([]models.Album, error) {
	// Execute the INSERT statement
	insertData := "INSERT INTO album (id, title, artist, price) VALUES ($1, $2, $3, $4)"
	for _, album := range a {
		_, err := dbConnection.Connection().Exec(insertData, album.ID, album.Title, album.Artist, album.Price)
		if err != nil {
			return nil, err
		}
	}
	// Return the newly inserted album
	return a, nil
}




//GetAlbumQuery returns a single album from the database
func GetAlbumQuery(id string) (*models.Album, error) {
	// Execute the SELECT statement
	var a models.Album
	err := dbConnection.Connection().QueryRow("SELECT id, title, artist, price FROM album WHERE id = $1", id).Scan(&a.ID, &a.Title, &a.Artist, &a.Price)
	if err != nil {
		return nil, err
	}
	// Return the album
	return &a, nil
}



// DeleteAlbumQuery removes an album from the database
func DeleteAlbumQuery(id string) error {
	// Execute the DELETE statement
	_, err := dbConnection.Connection().Exec("DELETE FROM album WHERE id = $1", id)
	if err != nil {
		return err
	}
	// Return any errors
	return nil
}


// PatchAlbumQuery updates an album in the database.
func PatchAlbumQuery(id string, a *models.Album) (*models.Album, error) {
	// Execute the UPDATE statement
	_, err := dbConnection.Connection().Exec("UPDATE album SET title = $1, artist = $2, price = $3 WHERE id = $4", a.Title, a.Artist, a.Price, id)
	if err != nil {
		return nil, err
	}
	// Return the updated album
	return a, nil
}




