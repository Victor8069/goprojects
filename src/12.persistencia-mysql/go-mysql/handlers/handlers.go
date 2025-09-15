package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

func ListContacts(db *sql.DB) {
	query := "SELECT * FROM CONTACT"

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println("\nLISTA DE CONTACTOS:")
	fmt.Println("---------------------------------------------------------------")
	for rows.Next() {
		contact := models.Contact{}

		var valueEmail sql.NullString
		err = rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "---"
		}

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("---------------------------------------------------------------")
	}
}

func GetContactByID(db *sql.DB, contactID int) {
	query := "SELECT * FROM CONTACT WHERE ID = ?"

	row := db.QueryRow(query, contactID)

	contact := models.Contact{}
	var valueEmail sql.NullString

	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontró ningún contacto con el ID %d", contactID)
		} else {
			log.Fatal(err)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "---"
	}
	fmt.Println("\nCONTACTO:")
	fmt.Println("---------------------------------------------------------------")

	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("---------------------------------------------------------------")

}

func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con éxito")
}

func UpdateContact(db *sql.DB, contact models.Contact) {
	query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE id = ?"

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto actualizado con éxito")
}

func DeleteContact(db *sql.DB, contactID int) {
	query := "DELETE FROM contact WHERE id = ?"

	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto eliminado con éxito")
}
