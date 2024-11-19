CREATE TABLE `user`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '사용자 식별자',
    `name` VARCHAR(20) NOT NULL COMMENT 'asdasdasdasd',
    `password` VARCHAR(20) NOT NULL COMMENT 'password hash',
    `role` VARCHAR(80) NOT NULL COMMENT 'user_role',
    `created` DATETIME(6) NOT NULL COMMENT 'record created time',
    `modified` DATETIME(6) NOT NULL COMMENT 'record modified time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uix_name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='사용자';

CREATE TABLE `task`
(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '테스크 식별자',
    `title` VARCHAR(120) NOT NULL COMMENT 'task_title',
    `status` VARCHAR(80) NOT NULL COMMENT 'task_status',
    `created` DATETIME(6) NOT NULL COMMENT 'record created time',
    `modified` DATETIME(6) NOT NULL COMMENT 'record modified time',   
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='테스크';