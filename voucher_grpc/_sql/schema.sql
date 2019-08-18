CREATE TABLE `voucher` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`code` VARCHAR(64) NOT NULL,
	`discount` FLOAT UNSIGNED NOT NULL,
	`start` TIMESTAMP NOT NULL DEFAULT '',
	`end` TIMESTAMP NOT NULL DEFAULT '0000-00-00 00:00:00',
	PRIMARY KEY (`id`),
	INDEX `code` (`code`)
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
AUTO_INCREMENT=1535
;

select v1.id, v1.code, v1.`start`, v1.`end`, v2.id, v2.`start`, v2.`end`
from voucher as v1
join voucher as v2
on v1.code = v2.code
where v2.id > v1.id
and  v2.`start`<= v1.`end` AND v2.`end` >= v1.`start` 
LIMIT 10;
