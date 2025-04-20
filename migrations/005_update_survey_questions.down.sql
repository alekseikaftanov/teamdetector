-- Clear new questions and options
TRUNCATE TABLE survey_questions CASCADE;
TRUNCATE TABLE survey_options CASCADE;

-- Restore original options
INSERT INTO survey_options (text, value) VALUES
    ('Strongly Disagree', 1),
    ('Disagree', 2),
    ('Neutral', 3),
    ('Agree', 4),
    ('Strongly Agree', 5);

-- Restore original questions
INSERT INTO survey_questions (text, category) VALUES
    ('How satisfied are you with the team''s communication?', 'Communication'),
    ('How well do you understand your role and responsibilities?', 'Role Clarity'),
    ('How comfortable are you with giving and receiving feedback?', 'Feedback'),
    ('How well does the team collaborate on projects?', 'Collaboration'),
    ('How satisfied are you with the team''s decision-making process?', 'Decision Making'); 