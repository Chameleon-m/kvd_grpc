CREATE TABLE
    IF NOT EXISTS book(
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
        `name` VARCHAR(255) NOT NULL
    ) ENGINE = INNODB;

CREATE TABLE
    IF NOT EXISTS author(
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
        `name` VARCHAR(255) NOT NULL
    ) ENGINE = INNODB;

CREATE TABLE
    IF NOT EXISTS book_author(
        `book_id` BIGINT UNSIGNED NOT NULL,
        `author_id` BIGINT UNSIGNED NOT NULL,
        PRIMARY KEY (`book_id`, `author_id`),
        INDEX (`book_id`),
        INDEX (`author_id`),
        CONSTRAINT `book_author_book_fk` FOREIGN KEY `book_fk` (`book_id`) REFERENCES `book` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
        CONSTRAINT `book_author_author_fk` FOREIGN KEY `author_fk` (`author_id`) REFERENCES `author` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    ) ENGINE = INNODB;