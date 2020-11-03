package db

func insert(stmt string) {
}

func remove() {

}

// AddCard adds a card to the card table
func AddCard(cc *CardComplete) {
	tx, err := conn.Begin()
	if err != nil {
		panic(err)
	}

	// Add card
	card, err := tx.Prepare("insert into cards(name, amount, comment) values(?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer card.Close()

	cardRes, err := card.Exec(cc.Name, cc.Amount, cc.Comment)
	if err != nil {
		panic(err)
	}

	cardID, err := cardRes.LastInsertId()

	// Add CardColors
	cardColor, err := tx.Prepare("insert into card_colors(card_id, color_id) values(?, ?)")
	if err != nil {
		panic(err)
	}
	defer cardColor.Close()

	for _, v := range cc.Colors {
		_, err = cardColor.Exec(cardID, v)
	}

	// Add CardTypes
	cardType, err := tx.Prepare("insert into card_types(card_id, type_id) values(?, ?)")
	if err != nil {
		panic(err)
	}
	defer cardType.Close()

	for _, v := range cc.Types {
		_, err = cardType.Exec(cardID, v)
	}

	// Add CardTypes
	cardSubtype, err := tx.Prepare("insert into card_subtypes(card_id, subtype_id) values(?, ?)")
	if err != nil {
		panic(err)
	}
	defer cardSubtype.Close()

	for _, v := range cc.Subtypes {
		_, err = cardSubtype.Exec(cardID, v)
	}

	// Add CardTypes
	cardSupertype, err := tx.Prepare("insert into card_supertypes(card_id, supertype_id) values(?, ?)")
	if err != nil {
		panic(err)
	}
	defer cardSupertype.Close()

	for _, v := range cc.Supertypes {
		_, err = cardSupertype.Exec(cardID, v)
	}

	tx.Commit()

}

// DeleteCard deletes a card from the card table
func DeleteCard() {

}

// AddColor adds a color to the color table
func AddColor() {

}

// DeleteColor TODO: Maybeeeeeeee...?
func DeleteColor() {}

// AddType adds a type to the type table
func AddType() {

}

// DeleteType TODO: Maybeeeeeeee...?
func DeleteType() {}

// AddSubtype adds a subtype to the subtype table
func AddSubtype() {

}

// DeleteSubtype deletes a subtype from the subtype table
func DeleteSubtype() {

}

// AddSupertype adds a supertype to the supertype table
func AddSupertype() {

}

// DeleteSupertype deletes a supertype from the supertype table
func DeleteSupertype() {

}
