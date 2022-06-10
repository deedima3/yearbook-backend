## Birthday Query

### Check User Birthday
```sql
SELECT * from user 
	WHERE MONTH(user.birthDate) = MONTH(CAST(NOW() as DATE)) 
	AND DAY(birthDate) = DAY(CAST(NOW() as DATE))
	AND user.userID = ?;
```

### Get This Week Birthday
```sql
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
```