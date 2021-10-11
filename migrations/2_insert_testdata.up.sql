INSERT INTO `questions` (`title`, `description`, `score`)
VALUES ('Test Question 1', 'Test Question 1 Description',1);

INSERT INTO `users` (`name`, `email`, `password`)
VALUES ('Foo','foo@foo.com','hogehoge');

INSERT INTO `teams` (`name`)
VALUES ('Team 1');

INSERT INTO `team_members` (`team_id`, `user_id`)
VALUES (1, 1);

INSERT INTO `team_opened_questions` (`team_id`, `question_id`)
VALUES (1, 1);


