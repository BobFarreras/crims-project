-- Reference schema for PocketBase collections

CREATE TABLE games (
  id TEXT PRIMARY KEY,
  code TEXT UNIQUE NOT NULL,
  state TEXT NOT NULL,
  seed TEXT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE players (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  user_id TEXT NOT NULL,
  role TEXT NOT NULL,
  status TEXT NOT NULL,
  is_host BOOLEAN NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE events (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  timestamp TEXT NOT NULL,
  location_id TEXT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE clues (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  type TEXT NOT NULL,
  state TEXT NOT NULL,
  reliability INTEGER NOT NULL,
  facts JSON,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE persons (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  name TEXT NOT NULL,
  official_story TEXT NOT NULL,
  truth_story TEXT NOT NULL,
  stress INTEGER NOT NULL,
  credibility INTEGER NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE hypotheses (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  title TEXT NOT NULL,
  strength_score INTEGER NOT NULL,
  status TEXT NOT NULL,
  node_ids JSON,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE accusations (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  player_id TEXT NOT NULL,
  suspect_id TEXT NOT NULL,
  motive_id TEXT NOT NULL,
  evidence_id TEXT NOT NULL,
  verdict TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE forensics (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  clue_id TEXT NOT NULL,
  result TEXT NOT NULL,
  confidence INTEGER NOT NULL,
  status TEXT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE timeline (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  timestamp TEXT NOT NULL,
  title TEXT NOT NULL,
  description TEXT NOT NULL,
  event_id TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE interrogations (
  id TEXT PRIMARY KEY,
  game_id TEXT NOT NULL,
  person_id TEXT NOT NULL,
  question TEXT NOT NULL,
  answer TEXT NOT NULL,
  tone TEXT NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);
