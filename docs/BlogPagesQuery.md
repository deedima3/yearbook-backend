## Blog Page Query

### Get All Pages
```sql
SELECT bp.blogID, bp.owner, bp.header_img, u.nickname, bp.description FROM blogpages bp
	JOIN user u ON bp.owner = u.userID;
```

### Get User Page
```sql
SELECT bp.blogID, bp.owner, bp.header_img, u.nickname, bp.description FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.userID = ?;
```

### Check User Page Exist
```sql
SELECT blogID FROM blogpages
	WHERE blogID = ?;
```

### Check Pages Exist
```sql
SELECT blogID FROM blogpages
```

### Check User Exist
```sql
SELECT owner FROM blogpages
	WHERE owner = ?;
```

### New Blog Page
```sql
INSERT INTO yearbook_db.blogpages(header_img,description,owner) VALUES(?,?,?);
```

### Search Blogpage
```sql
SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.nickname LIKE '%s%%'
	UNION
	SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.nim LIKE '%s%%';
```

### Count Search Res
```sql
SELECT COUNT(u.userID) FROM blogpages AS bp
	JOIN user as u ON bp.owner = u.userID
	WHERE u.nickname LIKE '%s%%' OR u.nim LIKE '%s%%';
```

### Search Nickname
```sql
SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.nickname LIKE '%s%%';
```

### Count Search Nickname
```sql
SELECT COUNT(u.userID) FROM blogpages AS bp
	JOIN user as u ON bp.owner = u.userID
	WHERE u.nickname LIKE '%s%%' OR u.nim LIKE '%s%%';
```

### Search Nim
```sql
SELECT u.userID, bp.header_img, bp.description, u.nickname, u.nim, u.image FROM blogpages bp
	JOIN user u ON bp.owner = u.userID
	WHERE u.nim LIKE '%s%%';
```

### Count Search NIM
```sql
SELECT COUNT(u.userID) FROM blogpages AS bp
	JOIN user as u ON bp.owner = u.userID
	WHERE u.nim LIKE '%s%%';
```

### Update Blog Pages
```sql
UPDATE yearbook_db.blogpages
	SET header_img=?,description=?
	WHERE blogID=?
```
