-- Create surveys table
CREATE TABLE IF NOT EXISTS surveys (
    id SERIAL PRIMARY KEY,
    team_id INTEGER NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_by INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create survey_questions table (predefined questions)
CREATE TABLE IF NOT EXISTS survey_questions (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    category VARCHAR(50) NOT NULL
);

-- Insert predefined questions
INSERT INTO survey_questions (text, category) VALUES
    ('How satisfied are you with the team''s communication?', 'Communication'),
    ('How well do you understand your role and responsibilities?', 'Role Clarity'),
    ('How comfortable are you with giving and receiving feedback?', 'Feedback'),
    ('How well does the team collaborate on projects?', 'Collaboration'),
    ('How satisfied are you with the team''s decision-making process?', 'Decision Making');

-- Create survey_options table (predefined options)
CREATE TABLE IF NOT EXISTS survey_options (
    id SERIAL PRIMARY KEY,
    text VARCHAR(255) NOT NULL,
    value INTEGER NOT NULL
);

-- Insert predefined options
INSERT INTO survey_options (text, value) VALUES
    ('Strongly Disagree', 1),
    ('Disagree', 2),
    ('Neutral', 3),
    ('Agree', 4),
    ('Strongly Agree', 5);

-- Create survey_responses table
CREATE TABLE IF NOT EXISTS survey_responses (
    id SERIAL PRIMARY KEY,
    survey_id INTEGER NOT NULL REFERENCES surveys(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    question_id INTEGER NOT NULL REFERENCES survey_questions(id) ON DELETE CASCADE,
    option_id INTEGER NOT NULL REFERENCES survey_options(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(survey_id, user_id, question_id) -- One response per user per question per survey
); 