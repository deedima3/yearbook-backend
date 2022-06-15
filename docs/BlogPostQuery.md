## Blog Post Query

### Insert New Post

```sql
INSERT INTO railway.blogpost
	(content, pages, upvote, downvote, title)
	VALUES(?, ?, ?, ?, ?);
```

### Delete Post

```sql
DELETE FROM blogpost
	WHERE postID = ?;
```

### Check Post Exist

```sql
SELECT postID FROM blogpost
	WHERE postID = ?;
```

### Select Upvote Downvote

```sql
SELECT upvote, downvote FROM blogpost
	WHERE postID = ?;
```

### Select Top 10 Twits

```sql
SELECT postID, title, content, upvote, downvote
	FROM railway.blogpost
    GROUP BY postID
    HAVING SUM(upvote) > 10
	ORDER BY upvote DESC
	LIMIT 10;
```

### Check Twits Exist

```sql
SELECT postID FROM blogpost;
```

### Update Upvote

```sql
UPDATE blogpost SET upvote = upvote + 1
	WHERE postID = %d
```

### Update Downvote

```sql
UPDATE blogpost SET downvote = downvote + 1
	WHERE postID = ?;
```
