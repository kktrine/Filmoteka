insert into film (1, title, description, date_premiere, rating)
                  values (1, 'Barbie', 'Barbie and Ken are having the time of their lives in the colorful and seemingly perfect world of Barbie Land. However, when they get a chance to go to the real world, they soon discover the joys and perils of living among humans.', 2023, 6.6),
                         (2, 'Twilight', 'When Bella Swan moves to a small town in the Pacific Northwest, she falls in love with Edward Cullen, a mysterious classmate who reveals himself to be a 108-year-old vampire.', 2008, 5.3),
                         (3, 'La La Land', 'While navigating their careers in Los Angeles, a pianist and an actress fall in love while attempting to reconcile their aspirations for the future.', 2016, 8.0);
insert into actor values
                      (1, 'Ryan Gosling', 'M', '1980-11-12'),
                      (2, 'Robert Pattinson', 'M', '1986-05-13'),
                      (3, 'Margot Robbie', 'F', '1990-07-02');
insert into actors_films values
                             (1, 1),
                             (1, 3),
                             (3, 1),
                             (2, 2);
insert into users values
                      (1, 'user1', 'password1', 'isUser'),
                      (2, 'user2', 'password2', 'isAdmin');