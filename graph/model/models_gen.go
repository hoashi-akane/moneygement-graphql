// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// ユーザ側アドバイザ利用登録
type AddAdviser struct {
	// userId
	UserID int `json:"userId"`
	// 家計簿ID
	LedgerID int `json:"ledgerId"`
	// アドバイザID
	AdviserID int `json:"adviserId"`
}

// アドバイザ一覧用入力
type AdviserListFilter struct {
	// first
	First int `json:"first"`
	// last
	Last int `json:"last"`
}

// アドバイザから見た会員リスト用
type AdviserMember struct {
	// 家計簿id
	ID int `json:"id"`
	// ユーザID
	UserID int `json:"userId"`
	// ニックネーム
	NickName string `json:"nickName"`
	// 家計簿名
	LedgerName string `json:"ledgerName"`
}

// カテゴリ
type Category struct {
	// id
	ID int `json:"Id"`
	// 名前
	Name string `json:"Name"`
	// 作成日
	CreatedAt string `json:"created_at"`
}

// チャット
type Chat struct {
	// id
	ID int `json:"id"`
	// 家計簿ID
	LedgerID int `json:"ledgerId"`
	// ユーザID
	UserID int `json:"userId"`
	// コメント
	Comment string `json:"comment"`
	// 作成日
	CreatedAt string `json:"createdAt"`
	// ニックネーム
	Nickname string `json:"nickname"`
}

// チャット一覧取得用フィルター
type ChatFilter struct {
	LedgerID int `json:"ledgerId"`
	First    int `json:"first"`
	Last     int `json:"last"`
}

// 家計簿削除用
type DeleteLedger struct {
	// 家計簿ID
	ID int `json:"id"`
	// ユーザID
	UserID int `json:"userId"`
}

// ユーザグループ間要素
type Enrollment struct {
	// id
	ID int `json:"id"`
	// ユーザid
	UserID int `json:"user_id"`
	// グループid
	GroupID int `json:"group_id"`
}

// 支出
type Expense struct {
	// ID
	ID int `json:"id"`
	// 家計簿ID
	LedgerID int `json:"ledgerId"`
	// カテゴリID
	CategoryID int `json:"categoryId"`
	// カテゴリ
	Category *Category `json:"category"`
	// 日付
	Date string `json:"date"`
	// 金額
	Amount int `json:"amount"`
	// メモ
	Note string `json:"note"`
}

// グループ
type Group struct {
	// id
	ID int `json:"id"`
	// 作成者
	Author int `json:"author"`
	// グループ名
	Name string `json:"name"`
}

// 収入
type Income struct {
	// ID
	ID int `json:"id"`
	// 家計簿ID
	LedgerID int `json:"ledgerId"`
	// カテゴリID
	CategoryID int `json:"categoryId"`
	// カテゴリ
	Category *Category `json:"category"`
	// 日付
	Date string `json:"date"`
	// 金額
	Amount int `json:"amount"`
	// メモ
	Note string `json:"note"`
}

// 家計簿
type Ledger struct {
	// ID
	ID int `json:"id"`
	// ユーザID
	UserID int `json:"userId"`
	// アドバイザIDあれば
	AdviserID int `json:"adviserId"`
	// グループIDあれば
	GroupID int `json:"groupId"`
	// 名前
	Name string `json:"name"`
	// 作成日
	CreatedAt string `json:"created_at"`
	// 収入リスト
	Incomes []*Income `json:"incomes"`
	// 支出リスト
	Expenses []*Expense `json:"expenses"`
}

// 家計簿関連
type LedgerEtc struct {
	// カテゴリ一覧
	CategoryList []*Category `json:"categoryList"`
	// 家計簿リスト取得
	Ledgers []*Ledger `json:"ledgers"`
	// 家計簿取得
	Ledger *Ledger `json:"ledger"`
	// 共有家計簿リスト取得
	ShareLedgers []*Ledger `json:"shareLedgers"`
	// アドバイザ側家計簿リスト
	AdviserLedgers []*Ledger `json:"adviserLedgers"`
}

// ログイン用入力
type LoginInfo struct {
	// メールアドレス
	Email string `json:"email"`
	// パスワード
	Password string `json:"password"`
	// FCMトークン
	Token string `json:"token"`
}

// アドバイザ作成用入力
type NewAdviser struct {
	// id
	ID int `json:"id"`
	// 本名
	Name string `json:"name"`
	// 説明
	Introduction string `json:"introduction"`
	// アドバイザ名
	AdviserName string `json:"adviserName"`
}

// チャット作成入力
type NewChat struct {
	// 家計簿ID
	LedgerID int `json:"ledgerId"`
	// ユーザID
	UserID int `json:"userId"`
	// コメント
	Comment string `json:"comment"`
}

// 家計簿の支出詳細の入力
type NewExpenseDetail struct {
	// 家計簿ID
	LedgerID int `json:"ledgerId"`
	// カテゴリID
	CategoryID int `json:"categoryId"`
	// 日付
	Date string `json:"date"`
	// 支出額
	Amount int `json:"amount"`
	// メモ
	Note string `json:"note"`
}

// グループ作成入力
type NewGroup struct {
	// ユーザID
	UserID int `json:"userId"`
	// グループ名
	GroupName string `json:"groupName"`
	// 家計簿名
	LedgerName string `json:"ledgerName"`
}

// グループ追加用入力
type NewGroupUser struct {
	// groupId
	GroupID int `json:"groupId"`
	// e-mail
	Email string `json:"email"`
	// ニックネーム
	NickName string `json:"nickName"`
}

// 家計簿の収入詳細の入力
type NewIncomeDetail struct {
	// 家計簿ID
	LedgerID int `json:"ledgerId"`
	// カテゴリID
	CategoryID int `json:"categoryId"`
	// 日付
	Date string `json:"date"`
	// 収入額
	Amount int `json:"amount"`
	// メモ
	Note string `json:"note"`
}

// 家計簿テーブル作成
type NewLedger struct {
	// ユーザID
	UserID int `json:"userId"`
	// 家計簿名
	Name string `json:"name"`
}

type NewSaving struct {
	UserID string `json:"userId"`
}

// 貯金詳細用入力
type NewSavingDetail struct {
	// 貯金ID
	SavingID int `json:"savingId"`
	// 貯金額
	SavingAmount int `json:"savingAmount"`
	// 貯金をした日付
	SavingDate string `json:"savingDate"`
	// メモ
	Note string `json:"note"`
}

// ユーザ作成用入力
type NewUser struct {
	// 名前
	Name string `json:"name"`
	// 別名
	NickName string `json:"nickName"`
	// メールアドレス
	Email string `json:"email"`
	// パスワード
	Password string `json:"password"`
}

// 貯金関連
type Saving struct {
	// 貯金取得
	SavingDetail *Savings `json:"savingDetail"`
	// 貯金額・収入額の和を返す
	SavingAmount *SavingAmountList `json:"savingAmount"`
	// 貯金詳細リスト取得
	SavingsDetails []*SavingsDetail `json:"savingsDetails"`
}

// 貯金額と収入額の和
type SavingAmountList struct {
	SavingAmount  int `json:"savingAmount"`
	ExpenseAmount int `json:"expenseAmount"`
}

// 貯金
type Savings struct {
	// id
	ID int `json:"id"`
	// ユーザID
	UserID int `json:"userId"`
	// 目標貯金額
	TargetAmount int `json:"targetAmount"`
}

// 貯金詳細
type SavingsDetail struct {
	// ID
	ID string `json:"id"`
	// 貯金ID
	SavingID int `json:"saving_id"`
	// 貯金金額
	SavingAmount int `json:"saving_amount"`
	// 貯金日
	SavingDate string `json:"saving_date"`
	// メモ
	Note string `json:"note"`
}

// ユーザ更新
type UpdateUser struct {
	// ユーザID
	ID int `json:"id"`
	// 本名
	Name string `json:"name"`
	// 別称
	Nickname string `json:"nickname"`
	// メールアドレス
	Email string `json:"email"`
	// アドバイザネーム
	AdviserName *string `json:"adviserName"`
	// 説明文
	Introduction *string `json:"introduction"`
	// 目標貯金額
	TargetAmount int `json:"targetAmount"`
}

// アドバイザから見た会員リスト用入力
type UseAdviserMemberFilter struct {
	// id
	UserID int `json:"userId"`
	// first
	First int `json:"first"`
	// last
	Last int `json:"last"`
}

// ユーザ関連
type User struct {
	// ユーザID
	ID int `json:"id"`
	// 本名
	Name string `json:"name"`
	// 別称
	Nickname string `json:"nickname"`
	// メールアドレス
	Email string `json:"email"`
	// アドバイザネーム
	AdviserName *string `json:"adviser_name"`
	// 説明文
	Introduction *string `json:"introduction"`
	// トークン
	Token string `json:"token"`
}

// ユーザ認証情報
type UserAuth struct {
	// ユーザID
	UserID int `json:"userId"`
	// パスワード
	Password string `json:"password"`
}

// 貯金詳細取得フィルター
type SavingsDetailsFilter struct {
	// 貯金ID
	SavingsID int `json:"savings_id"`
	// 取得開始列
	First int `json:"first"`
	// 取得終了列
	Last int `json:"last"`
}
