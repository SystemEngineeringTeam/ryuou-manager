INSERT INTO `questions` ( `title`, `description`)
VALUES ('Test Question 1', 'Test Question 1 Description');

INSERT INTO `users` (`id`, `name`, `email`, `password`)
VALUES (`1`,`Foo`,`foo@foo.com`,`hogehoge`);

INSERT INTO `teams` (`name`)
VALUES ('Team 1');

INSERT INTO `team_members` (`team_id`, `user_id`)
VALUES (1, 1);

INSERT INTO `team_open_questions` (`team_id`, `question_id`)
VALUES (1, 1);


