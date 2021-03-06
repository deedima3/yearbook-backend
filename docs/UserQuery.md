## User Query

### Insert New User

```sql
INSERT INTO railway.user(email,password,nickname,nim) VALUES(?,?,?,?)
```

### Select User Per ID

```sql
SELECT userID FROM railway.user WHERE userID = ?
```

### Update User

```sql
UPDATE railway.user SET email = ?, password = ?, image = ?, nickname = ? WHERE userID = ?
```

### Get All User

```sql
SELECT userID, email, nickname FROM railway.user
```

### Get User Password

```sql
SELECT userID, email, nickname, password FROM railway.user WHERE email = ?
```

### Trigger Backup User

```sql
DELIMITER $$
DROP TRIGGER IF EXISTS log_user;
CREATE TRIGGER log_user
    AFTER DELETE ON user
    BEGIN
        INSERT INTO user_log(email, password, nickname)
        VALUES(OLD.email, OLD.password, OLD.nickname);
    END$$
DELIMITER ;
```
