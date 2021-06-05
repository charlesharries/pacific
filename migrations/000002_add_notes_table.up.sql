CREATE TABLE `notes` (
  `id` integer,
  `created_at` datetime,
  `updated_at` datetime,
  `deleted_at` datetime,
  `user_id` integer,
  `date` datetime,
  `content` text,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_notes_user` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);
CREATE UNIQUE INDEX notes_date_user_id_idx ON notes(user_id, date);