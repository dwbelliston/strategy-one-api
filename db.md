# DB

- connect

Local tunnel
```bash
psql --host=127.0.0.1 --port=5432 --username=awsuser --password --dbname=dev
```

psql
```bash
psql \
  --host=strategy-one-db.cluster-ciwkcai1iw95.us-east-1.rds.amazonaws.com \
  --port=5432 \
  --username=awsuser \
  --password \
  --dbname=dev
```

mysql
```
mysql --user=awsuser --password -h strategy-one-db.cluster-ciwkcai1iw95.us-east-1.rds.amazonaws.com
```

- create table

```redshift
create table shapes(
  shape_id bigint identity(0, 1),
  sides integer not null,
  title varchar(30) not null,
  created_at datetime default sysdate
);
```

```psql
CREATE TABLE shapes(
 shape_id serial PRIMARY KEY,
 title VARCHAR (355) NOT NULL,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 sides integer NOT NULL
);
```

```mysql
CREATE TABLE IF NOT EXISTS shapes (
    shape_id INT AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    sides TINYINT NOT NULL,
    PRIMARY KEY (shape_id)
);
```


- add records

INSERT INTO shapes (sides, title)
VALUES (0, 'Circle'), (4, 'Square');
