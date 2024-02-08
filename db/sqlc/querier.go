// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	CreateFund(ctx context.Context, arg CreateFundParams) (GlFund, error)
	CreateGLAccount(ctx context.Context, arg CreateGLAccountParams) (GlAccount, error)
	CreateJournalDetails(ctx context.Context, arg CreateJournalDetailsParams) (GlJournalDetail, error)
	CreateJournalHeader(ctx context.Context, arg CreateJournalHeaderParams) (GlJournalHeader, error)
	CreatePeriod(ctx context.Context, arg CreatePeriodParams) (GlPeriod, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateTask(ctx context.Context, arg CreateTaskParams) (KbTask, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerifyEmail(ctx context.Context, arg CreateVerifyEmailParams) (VerifyEmail, error)
	GetDistLedgerByRef(ctx context.Context, arg GetDistLedgerByRefParams) (GlDistributionLedger, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetFund(ctx context.Context, fund string) (GlFund, error)
	GetGLAccount(ctx context.Context, arg GetGLAccountParams) (GlAccount, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetTaskList(ctx context.Context) ([]GetTaskListRow, error)
	GetTasks(ctx context.Context, taskID string) ([]KbTask, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListAccountTypes(ctx context.Context) ([]GlAccountType, error)
	ListDistLedger(ctx context.Context) ([]GlDistributionLedger, error)
	ListDistLedgerByPeriod(ctx context.Context, period int32) ([]GlDistributionLedger, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListFunds(ctx context.Context) ([]GlFund, error)
	ListGeneralLedger(ctx context.Context) ([]GlAccount, error)
	ListJournalDetails(ctx context.Context) ([]GlJournalDetail, error)
	ListJournalHeader(ctx context.Context) ([]GlJournalHeader, error)
	ListPeriods(ctx context.Context) ([]GlPeriod, error)
	ListPriority(ctx context.Context) ([]KbPriority, error)
	ListStatus(ctx context.Context) ([]KbStatus, error)
	ListSubTask(ctx context.Context) ([]KbSubtask, error)
	ListTasks(ctx context.Context) ([]KbTask, error)
	ListTeam(ctx context.Context) ([]KbTeam, error)
	ListType(ctx context.Context) ([]KbType, error)
	UpdateFund(ctx context.Context, arg UpdateFundParams) (GlFund, error)
	UpdateGLAccount(ctx context.Context, arg UpdateGLAccountParams) (GlAccount, error)
	UpdateJournalDetails(ctx context.Context, arg UpdateJournalDetailsParams) (GlJournalDetail, error)
	UpdateJournalHeader(ctx context.Context, arg UpdateJournalHeaderParams) (GlJournalHeader, error)
	UpdatePeriod(ctx context.Context, arg UpdatePeriodParams) (GlPeriod, error)
	UpdateTask(ctx context.Context, arg UpdateTaskParams) (KbTask, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateVerifyEmail(ctx context.Context, arg UpdateVerifyEmailParams) (VerifyEmail, error)
}

var _ Querier = (*Queries)(nil)
