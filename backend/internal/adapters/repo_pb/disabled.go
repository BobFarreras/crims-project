package repo_pb

import (
	"context"

	"github.com/digitaistudios/crims-backend/internal/ports"
)

// DisabledPocketBaseClient implementa un client que sempre retorna error.
type DisabledPocketBaseClient struct {
	Err error
}

func (d DisabledPocketBaseClient) Ping(ctx context.Context) error {
	return d.Err
}

type DisabledGameRepository struct {
	Err error
}

func (d DisabledGameRepository) CreateGame(ctx context.Context, input ports.GameRecordInput) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.Err
}

func (d DisabledGameRepository) GetGameByID(ctx context.Context, id string) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.Err
}

func (d DisabledGameRepository) GetGameByCode(ctx context.Context, code string) (ports.GameRecord, error) {
	return ports.GameRecord{}, d.Err
}

type DisabledPlayerRepository struct {
	Err error
}

func (d DisabledPlayerRepository) CreatePlayer(ctx context.Context, input ports.PlayerRecordInput) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, d.Err
}

func (d DisabledPlayerRepository) GetPlayerByID(ctx context.Context, id string) (ports.PlayerRecord, error) {
	return ports.PlayerRecord{}, d.Err
}

func (d DisabledPlayerRepository) ListPlayersByGame(ctx context.Context, gameID string) ([]ports.PlayerRecord, error) {
	return nil, d.Err
}

type DisabledEventRepository struct {
	Err error
}

func (d DisabledEventRepository) CreateEvent(ctx context.Context, input ports.EventRecordInput) (ports.EventRecord, error) {
	return ports.EventRecord{}, d.Err
}

func (d DisabledEventRepository) GetEventByID(ctx context.Context, id string) (ports.EventRecord, error) {
	return ports.EventRecord{}, d.Err
}

func (d DisabledEventRepository) ListEventsByGame(ctx context.Context, gameID string) ([]ports.EventRecord, error) {
	return nil, d.Err
}

type DisabledClueRepository struct {
	Err error
}

func (d DisabledClueRepository) CreateClue(ctx context.Context, input ports.ClueRecordInput) (ports.ClueRecord, error) {
	return ports.ClueRecord{}, d.Err
}

func (d DisabledClueRepository) GetClueByID(ctx context.Context, id string) (ports.ClueRecord, error) {
	return ports.ClueRecord{}, d.Err
}

func (d DisabledClueRepository) ListCluesByGame(ctx context.Context, gameID string) ([]ports.ClueRecord, error) {
	return nil, d.Err
}

type DisabledPersonRepository struct {
	Err error
}

func (d DisabledPersonRepository) CreatePerson(ctx context.Context, input ports.PersonRecordInput) (ports.PersonRecord, error) {
	return ports.PersonRecord{}, d.Err
}

func (d DisabledPersonRepository) GetPersonByID(ctx context.Context, id string) (ports.PersonRecord, error) {
	return ports.PersonRecord{}, d.Err
}

func (d DisabledPersonRepository) ListPersonsByGame(ctx context.Context, gameID string) ([]ports.PersonRecord, error) {
	return nil, d.Err
}

type DisabledHypothesisRepository struct {
	Err error
}

func (d DisabledHypothesisRepository) CreateHypothesis(ctx context.Context, input ports.HypothesisRecordInput) (ports.HypothesisRecord, error) {
	return ports.HypothesisRecord{}, d.Err
}

func (d DisabledHypothesisRepository) GetHypothesisByID(ctx context.Context, id string) (ports.HypothesisRecord, error) {
	return ports.HypothesisRecord{}, d.Err
}

func (d DisabledHypothesisRepository) ListHypothesesByGame(ctx context.Context, gameID string) ([]ports.HypothesisRecord, error) {
	return nil, d.Err
}

type DisabledAccusationRepository struct {
	Err error
}

func (d DisabledAccusationRepository) CreateAccusation(ctx context.Context, input ports.AccusationRecordInput) (ports.AccusationRecord, error) {
	return ports.AccusationRecord{}, d.Err
}

func (d DisabledAccusationRepository) GetAccusationByID(ctx context.Context, id string) (ports.AccusationRecord, error) {
	return ports.AccusationRecord{}, d.Err
}

func (d DisabledAccusationRepository) ListAccusationsByGame(ctx context.Context, gameID string) ([]ports.AccusationRecord, error) {
	return nil, d.Err
}

type DisabledForensicRepository struct {
	Err error
}

func (d DisabledForensicRepository) CreateAnalysis(ctx context.Context, input ports.ForensicRecordInput) (ports.ForensicRecord, error) {
	return ports.ForensicRecord{}, d.Err
}

func (d DisabledForensicRepository) GetAnalysisByID(ctx context.Context, id string) (ports.ForensicRecord, error) {
	return ports.ForensicRecord{}, d.Err
}

func (d DisabledForensicRepository) ListAnalysesByGame(ctx context.Context, gameID string) ([]ports.ForensicRecord, error) {
	return nil, d.Err
}

type DisabledTimelineRepository struct {
	Err error
}

func (d DisabledTimelineRepository) CreateEntry(ctx context.Context, input ports.TimelineRecordInput) (ports.TimelineRecord, error) {
	return ports.TimelineRecord{}, d.Err
}

func (d DisabledTimelineRepository) GetEntryByID(ctx context.Context, id string) (ports.TimelineRecord, error) {
	return ports.TimelineRecord{}, d.Err
}

func (d DisabledTimelineRepository) ListEntriesByGame(ctx context.Context, gameID string) ([]ports.TimelineRecord, error) {
	return nil, d.Err
}

type DisabledInterrogationRepository struct {
	Err error
}

func (d DisabledInterrogationRepository) CreateInterrogation(ctx context.Context, input ports.InterrogationRecordInput) (ports.InterrogationRecord, error) {
	return ports.InterrogationRecord{}, d.Err
}

func (d DisabledInterrogationRepository) GetInterrogationByID(ctx context.Context, id string) (ports.InterrogationRecord, error) {
	return ports.InterrogationRecord{}, d.Err
}

func (d DisabledInterrogationRepository) ListInterrogationsByGame(ctx context.Context, gameID string) ([]ports.InterrogationRecord, error) {
	return nil, d.Err
}

// AFEGEIX AQUEST MÈTODE AL FINAL DE TOT
func (d DisabledPocketBaseClient) CreateUser(username, email, password, passwordConfirm, name string) error {
	return d.Err
}

// Afegeix aquest mètode a disabled.go
func (d DisabledPocketBaseClient) AuthWithPassword(identity, password string) (*ports.AuthResponse, error) {
	return nil, d.Err
}
