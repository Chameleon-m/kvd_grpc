USE `library`;

INSERT author
VALUES (1, 'Author 1'), (2, 'Author 2'), (3, 'Author 3'), (4, 'Author 4'), (5, 'Author 5'), (6, 'Author 6'), (7, 'Author 7'), (8, 'Author 8'), (9, 'Author 9'), (10, 'Author 10'), (11, 'Author 11'), (12, 'Author 12'), (13, 'Author 13'), (14, 'Author 14'), (15, 'Author 15'), (16, 'Author 16'), (17, 'Author 17'), (18, 'Author 18'), (19, 'Author 19'), (20, 'Author 20'), (21, 'Author 21'), (22, 'Author 22'), (23, 'Author 23'), (24, 'Author 24'), (25, 'Author 25'), (26, 'Author 26'), (27, 'Author 27'), (28, 'Author 28'), (29, 'Author 29'), (30, 'Author 30'), (31, 'Author 31'), (32, 'Author 32'), (33, 'Author 33'), (34, 'Author 34'), (35, 'Author 35'), (36, 'Author 36'), (37, 'Author 37'), (38, 'Author 38'), (39, 'Author 39'), (40, 'Author 40'), (41, 'Author 41'), (42, 'Author 42'), (43, 'Author 43'), (44, 'Author 44'), (45, 'Author 45'), (46, 'Author 46'), (47, 'Author 47'), (48, 'Author 48'), (49, 'Author 49'), (50, 'Author 50'), (51, 'Author 51'), (52, 'Author 52'), (53, 'Author 53'), (54, 'Author 54'), (55, 'Author 55'), (56, 'Author 56'), (57, 'Author 57'), (58, 'Author 58'), (59, 'Author 59'), (60, 'Author 60'), (61, 'Author 61'), (62, 'Author 62'), (63, 'Author 63'), (64, 'Author 64'), (65, 'Author 65'), (66, 'Author 66'), (67, 'Author 67'), (68, 'Author 68'), (69, 'Author 69'), (70, 'Author 70'), (71, 'Author 71'), (72, 'Author 72'), (73, 'Author 73'), (74, 'Author 74'), (75, 'Author 75'), (76, 'Author 76'), (77, 'Author 77'), (78, 'Author 78'), (79, 'Author 79'), (80, 'Author 80'), (81, 'Author 81'), (82, 'Author 82'), (83, 'Author 83'), (84, 'Author 84'), (85, 'Author 85'), (86, 'Author 86'), (87, 'Author 87'), (88, 'Author 88'), (89, 'Author 89'), (90, 'Author 90'), (91, 'Author 91'), (92, 'Author 92'), (93, 'Author 93'), (94, 'Author 94'), (95, 'Author 95'), (96, 'Author 96'), (97, 'Author 97'), (98, 'Author 98'), (99, 'Author 99'), (100, 'Author 100');

INSERT book
VALUES (1, 'Book 1'), (2, 'Book 2'), (3, 'Book 3'), (4, 'Book 4'), (5, 'Book 5'), (6, 'Book 6'), (7, 'Book 7'), (8, 'Book 8'), (9, 'Book 9'), (10, 'Book 10'), (11, 'Book 11'), (12, 'Book 12'), (13, 'Book 13'), (14, 'Book 14'), (15, 'Book 15'), (16, 'Book 16'), (17, 'Book 17'), (18, 'Book 18'), (19, 'Book 19'), (20, 'Book 20'), (21, 'Book 21'), (22, 'Book 22'), (23, 'Book 23'), (24, 'Book 24'), (25, 'Book 25'), (26, 'Book 26'), (27, 'Book 27'), (28, 'Book 28'), (29, 'Book 29'), (30, 'Book 30'), (31, 'Book 31'), (32, 'Book 32'), (33, 'Book 33'), (34, 'Book 34'), (35, 'Book 35'), (36, 'Book 36'), (37, 'Book 37'), (38, 'Book 38'), (39, 'Book 39'), (40, 'Book 40'), (41, 'Book 41'), (42, 'Book 42'), (43, 'Book 43'), (44, 'Book 44'), (45, 'Book 45'), (46, 'Book 46'), (47, 'Book 47'), (48, 'Book 48'), (49, 'Book 49'), (50, 'Book 50'), (51, 'Book 51'), (52, 'Book 52'), (53, 'Book 53'), (54, 'Book 54'), (55, 'Book 55'), (56, 'Book 56'), (57, 'Book 57'), (58, 'Book 58'), (59, 'Book 59'), (60, 'Book 60'), (61, 'Book 61'), (62, 'Book 62'), (63, 'Book 63'), (64, 'Book 64'), (65, 'Book 65'), (66, 'Book 66'), (67, 'Book 67'), (68, 'Book 68'), (69, 'Book 69'), (70, 'Book 70'), (71, 'Book 71'), (72, 'Book 72'), (73, 'Book 73'), (74, 'Book 74'), (75, 'Book 75'), (76, 'Book 76'), (77, 'Book 77'), (78, 'Book 78'), (79, 'Book 79'), (80, 'Book 80'), (81, 'Book 81'), (82, 'Book 82'), (83, 'Book 83'), (84, 'Book 84'), (85, 'Book 85'), (86, 'Book 86'), (87, 'Book 87'), (88, 'Book 88'), (89, 'Book 89'), (90, 'Book 90'), (91, 'Book 91'), (92, 'Book 92'), (93, 'Book 93'), (94, 'Book 94'), (95, 'Book 95'), (96, 'Book 96'), (97, 'Book 97'), (98, 'Book 98'), (99, 'Book 99'), (100, 'Book 100');

INSERT book_author
VALUES (1, 1), (1, 2), (3, 3), (3, 4), (5, 5), (5, 6), (7, 7), (7, 8), (9, 9), (9, 10), (11, 11), (11, 12), (13, 13), (13, 14), (15, 15), (15, 16), (17, 17), (17, 18), (19, 19), (19, 20), (21, 21), (21, 22), (23, 23), (23, 24), (25, 25), (25, 26), (27, 27), (27, 28), (29, 29), (29, 30), (31, 31), (31, 32), (33, 33), (33, 34), (35, 35), (35, 36), (37, 37), (37, 38), (39, 39), (39, 40), (41, 41), (41, 42), (43, 43), (43, 44), (45, 45), (45, 46), (47, 47), (47, 48), (49, 49), (49, 50), (51, 51), (51, 52), (53, 53), (53, 54), (55, 55), (55, 56), (57, 57), (57, 58), (59, 59), (59, 60), (61, 61), (61, 62), (63, 63), (63, 64), (65, 65), (65, 66), (67, 67), (67, 68), (69, 69), (69, 70), (71, 71), (71, 72), (73, 73), (73, 74), (75, 75), (75, 76), (77, 77), (77, 78), (79, 79), (79, 80), (81, 81), (81, 82), (83, 83), (83, 84), (85, 85), (85, 86), (87, 87), (87, 88), (89, 89), (89, 90), (91, 91), (91, 92), (93, 93), (93, 94), (95, 95), (95, 96), (97, 97), (97, 98), (99, 99), (99, 100);