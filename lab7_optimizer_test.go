package main

import "testing"

func TestRunLab7OptimizerBuildsStructuredResult(t *testing.T) {
	app := NewApp()
	result := app.RunLab7Optimizer("fun calc(a: Int, b: Int, c: Int): Int { return a + b * c };")

	if result.OutputKey != "lab7.success" {
		t.Fatalf("expected lab7.success, got output key %q", result.OutputKey)
	}
	if len(result.Errors) != 0 {
		t.Fatalf("expected no errors, got %+v", result.Errors)
	}
	if result.Output != "" {
		t.Fatalf("expected structured params instead of plain output, got %q", result.Output)
	}

	params := result.OutputParams
	if params["ast"] == "" {
		t.Fatal("expected AST in output params")
	}
	if params["inputIR"] != "1. t1 = b * c\n2. t2 = a + t1\n3. return t2" {
		t.Fatalf("unexpected input IR:\n%s", params["inputIR"])
	}
	if params["foldChanged"] != "false" || params["neutralChanged"] != "false" {
		t.Fatalf("expected calc expression to stay unchanged, got fold=%s neutral=%s", params["foldChanged"], params["neutralChanged"])
	}
	if params["optimizedOutputIR"] != params["inputIR"] {
		t.Fatalf("expected final IR to equal input IR for calc")
	}
}

func TestLab7ConstantFolding(t *testing.T) {
	program, failed := buildLab7Program("fun constCalc(): Int { return 2 + 3 * 4 };")
	if failed != nil {
		t.Fatalf("unexpected parse failure: %+v", failed)
	}

	folded, changed := lab7FoldConstants(program.expr)
	if !changed {
		t.Fatal("expected constant folding to change expression")
	}

	got := lab7FormatTAC(lab7GenerateTAC(folded))
	if got != "1. return 14" {
		t.Fatalf("unexpected folded TAC:\n%s", got)
	}
}

func TestLab7NeutralOperationSimplification(t *testing.T) {
	program, failed := buildLab7Program("fun identityCalc(a: Int, b: Int): Int { return (a + 0) + (b * 1) };")
	if failed != nil {
		t.Fatalf("unexpected parse failure: %+v", failed)
	}

	simplified, changed := lab7SimplifyNeutral(program.expr)
	if !changed {
		t.Fatal("expected neutral operation simplification to change expression")
	}

	got := lab7FormatTAC(lab7GenerateTAC(simplified))
	want := "1. t1 = a + b\n2. return t1"
	if got != want {
		t.Fatalf("unexpected simplified TAC:\nwant:\n%s\ngot:\n%s", want, got)
	}
}

func TestRunLab7OptimizerReportsSyntaxErrors(t *testing.T) {
	app := NewApp()
	result := app.RunLab7Optimizer("fun calc(a: Int): Int { return a + };")

	if result.OutputKey != "lab7.parserFailed" {
		t.Fatalf("expected lab7.parserFailed, got %q", result.OutputKey)
	}
	if len(result.Errors) == 0 {
		t.Fatal("expected parser errors")
	}
}
