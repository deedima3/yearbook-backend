package impl

import (
	"context"
	"database/sql"
	"log"

	"github.com/deedima3/yearbook-backend/internal/blogpages/entity"
)

type birthdayRepositoryImpl struct {
	DB *sql.DB
}

var (
	GET_ALL_WEEK_BIRTHDAY = `
	SELECT
bp.blogID,
bp.owner,
bp.header_img,
u.nickname,
bp.description,
u.birthDate
FROM
    blogpages bp
    JOIN user u ON bp.owner = u.userID
WHERE
    month(birthdate) BETWEEN month(NOW() - INTERVAL 3 day)
    AND month(NOW() + INTERVAL 3 day)
    AND day(u.birthDate) BETWEEN day(NOW() - INTERVAL 3 day)
    AND day(NOW() + INTERVAL 3 day);
	`
	CHECK_USER_BIRTHDAY = `
	SELECT * from user 
	WHERE MONTH(user.birthDate) = MONTH(CAST(NOW() as DATE)) 
	AND DAY(birthDate) = DAY(CAST(NOW() as DATE))
	AND user.userID = ?
	;
	`
)

func ProvideBirthdayRepository(DB *sql.DB) *birthdayRepositoryImpl {
	return &birthdayRepositoryImpl{DB: DB}
}

func (br birthdayRepositoryImpl) CheckUserBirthday(ctx context.Context, id uint64) (bool, error) {
	query := CHECK_USER_BIRTHDAY
	stmt, err := br.DB.PrepareContext(ctx, query)

	if err != nil {
		log.Printf("ERROR CheckPostExists -> error: %v\n", err)
		return false, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		log.Printf("ERROR CheckPostExists -> error: %v\n", err)
		return false, err
	}
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (br birthdayRepositoryImpl) GetBirthdayWeek(ctx context.Context) (entity.BlogPagesPeopleJoined, error) {
	query := GET_ALL_WEEK_BIRTHDAY

	stmt, err := br.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("ERROR GetAllPages -> error: %v\n", err)
		return entity.BlogPagesPeopleJoined{}, err
	}
	rows, err := stmt.Query()
	if err != nil {
		log.Printf("ERROR GetAllPages -> error: %v\n", err)
		return entity.BlogPagesPeopleJoined{}, err
	}

	blogPages := entity.BlogPagesPeopleJoined{}

	for rows.Next() {
		var blogPage entity.BlogPageUserJoined

		err := rows.Scan(
			&blogPage.BlogPage.PageID,
			&blogPage.BlogPage.Owner,
			&blogPage.BlogPage.HeaderImage,
			&blogPage.User.Nickname,
			&blogPage.BlogPage.Description,
			&blogPage.User.BirthDate,
		)

		if err != nil {
			log.Printf("ERROR GetAllPages -> error: %v\n", err)
			return entity.BlogPagesPeopleJoined{}, err
		}
		blogPages = append(blogPages, &blogPage)
	}
	return blogPages, nil
}
