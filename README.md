## RideCell Assignent

### QuickStart
In your GOPATH/github.com/bluehawk27
    
    $ git clone https://github.com/bluehawk27/ridecell.git
    $ mysql -u root

Create the DB and add Some dummy Data

    CREATE DATABASE ridecell CHARACTER SET utf8;
    CREATE TABLE `parkingSpot` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `available` tinyint(1) NOT NULL DEFAULT '0',
    `lat` double DEFAULT NULL,
    `long` double DEFAULT NULL,
    `reservedAt` varchar(255) DEFAULT NULL,
    `price` float DEFAULT NULL,
    `startTime` varchar(255) DEFAULT NULL,
    `endTime` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `lat_long_idx` (`lat`,`long`)
    )   ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;


    INSERT INTO `parkingSpot` (`id`, `available`, `lat`, `long`, `reservedAt`, `price`, `startTime`, `endTime`)
    VALUES
	(1,0,37.781319,-122.395962,NULL,34.99,NULL,NULL),
	(2,0,37.781141,-122.395715,NULL,NULL,NULL,NULL),
	(3,0,37.78279,-122.3972,NULL,NULL,NULL,NULL),
	(4,0,37.784872,-122.397549,NULL,NULL,NULL,NULL),
	(5,0,37.782568,-122.433598,NULL,NULL,NULL,NULL),
	(6,0,37.48654,-122.226354,NULL,NULL,NULL,NULL),
	(7,0,37.487281,-122.226418,NULL,NULL,NULL,NULL),
	(8,0,37.489194,-122.229346,NULL,NULL,NULL,NULL),
	(9,0,37.490233,-122.233257,NULL,NULL,NULL,NULL),
	(10,0,37.491442,-122.234866,NULL,NULL,NULL,NULL);


### Tests

    make test

#### Start the Service

    make start

#### Sample request

http://127.0.0.1:8080/list/?lat=37.486714&long=-122.226306&radius=.5
