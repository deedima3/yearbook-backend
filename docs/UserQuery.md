## User Query

### Insert New User
```sql
INSERT INTO yearbook_db.user(email,password,nickname,nim) VALUES(?,?,?,?)
```

### Select User Per ID
```sql
SELECT userID FROM yearbook_db.user WHERE userID = ?
```

### Update User
```sql
UPDATE yearbook_db.user SET email = ?, password = ?, image = ?, nickname = ? WHERE userID = ?
 ```

### Get All User
```sql
SELECT userID, email, nickname FROM yearbook_db.user
```

### Get User Password
```sql
SELECT userID, email, nickname, password FROM yearbook_db.user WHERE email = ?
```
