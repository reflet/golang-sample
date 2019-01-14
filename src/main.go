package main

import (
	"pkg1"
	"workpkg/workhello"
	"workpkg/workvar"
	"workpkg/worklit"
	"workpkg/workconst"
	"workpkg/workopt"
	"workpkg/workcast"
	"workpkg/workif"
	"workpkg/workswitch"
	"workpkg/workfor"
	"workpkg/workgoto"
	"workpkg/workfloat"
)

/**
 * main関数
 */
func main() {
	/**
	 * 複数パッケージの構成 ＋ vendoringの挙動確認
	 *「workhello」パッケージの「Hello」関数を使用する
	 */
	workhello.Hello()

	/**
	 * 相対パスでimportする例
	 * ※「pkg1」パッケージの「F1」関数を使用する
	 */
	pkg1.F1()

	/**
	 * 変数の宣言
	 */
	workvar.Variable()

	/**
     * 変数の宣言(複数)
     */
	workvar.Variable1()

	/**
     * 変数の宣言(複数 - 関数の戻り値あり)
     */
	workvar.Variable2()

	/**
     * 変数を宣言(:=)
     */
	workvar.Variable3()

	/**
	 * グローバル変数
	 */
	workvar.Global()

	/**
     * ローカル変数
     */
	workvar.LocalVariable()

	/**
	 * リテラル (整数)
	 */
	worklit.Literal()

	/**
	 * リテラル (文字列)
	 */
	worklit.LiteralStr()

	/**
	 * 定数の宣言
	 */
	workconst.Constunt()

	/**
	 * iota(イオタ)
	 */
	workconst.IotaConstunt()

	/**
	 * 演算子
	 */
	workopt.Operator()

	/**
     * 比較演算子
     */
	workopt.ComparisonOperator()

	/**
     * 論理演算子
     */
	workopt.LogicalOperator()

	/**
     * インクリメント/デクリメント
     */
	workopt.IncrementDecrement()

	/**
     * 複素数
     */
	workopt.ComplexCalc()

	/**
     * 文字列の操作
     */
	workopt.StrOperation()

	/**
	 * 型変換（キャスト）
	 */
	workcast.Cast()

	/**
	 * 定数式
	 */
	workconst.ConstantExpression()

	/**
	 * 条件分岐(if文)
	 */
	workif.If()

	/**
	 * 条件分岐(switch文)
	 */
	workswitch.Switch()

	/**
	 * 条件分岐(switch文 - fallthrough文)
	 */
	workswitch.Fallthrough()

	/**
	 * 条件分岐(switch文 - break文)
	 */
	workswitch.Break()

	/**
	 * 条件分岐(for文)
	 */
	workfor.For()

	/**
	 * 条件分岐(for文 - continue文 / break文)
	 */
	workfor.ContinueBreak()

	/**
     * 条件分岐(ラベル付きのfor文)
     */
	workfor.Label()

	/**
     * 条件分岐(rangeを用いたfor文)
	 */
	workfor.Range()

	/**
	 * ラベルへのジャンプ(goto文)
	 */
	workgoto.Goto()

	/**
	 *
	 */
	workfloat.Float()
}
