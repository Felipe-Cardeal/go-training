package main

import "testing"

func Test_adicionarPercentual(t *testing.T) {
	type args struct {
		valor       float64
		porcentagem float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adicionarPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
				t.Errorf("adicionarPercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_descobrirPercentual(t *testing.T) {
	type args struct {
		valor       float64
		porcentagem float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
		}
		for _, tt := range tests{
		t.Run(tt.name, func (t *testing.T){
		if got := descobrirPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want{
		t.Errorf("descobrirPercentual() = %v, want %v", got, tt.want)
	}
	})
	}
	}

	func
	Test_divisao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := divisao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("divisao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_multiplicacao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := multiplicacao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_soma(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := soma(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("soma() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_subtracao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want int
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("subtracao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_adicionarPercentual1(t * testing.T)
	{
		type args struct {
			valor       float64
			porcentagem float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := adicionarPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
					t.Errorf("adicionarPercentual() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_descobrirPercentual1(t * testing.T)
	{
		type args struct {
			valor       float64
			porcentagem float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := descobrirPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
					t.Errorf("descobrirPercentual() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_divisao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := divisao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("divisao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_multiplicacao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := multiplicacao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_soma(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := soma(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("soma() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func Test_subtracao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}		}{

	name:"teste"2 + 1,
		args: args{x2, y1},
		want: 3,
	}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("subtracao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_adicionarPercentual1(t * testing.T)
	{
		type args struct {
			valor       float64
			porcentagem float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := adicionarPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
					t.Errorf("adicionarPercentual() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_descobrirPercentual1(t * testing.T)
	{
		type args struct {
			valor       float64
			porcentagem float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := descobrirPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
					t.Errorf("descobrirPercentual() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_divisao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := divisao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("divisao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_multiplicacao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := multiplicacao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_soma(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := soma(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("soma() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_subtracao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("subtracao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_adicionarPercentual1(t * testing.T)
	{
		type args struct {
			valor       float64
			porcentagem float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := adicionarPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
					t.Errorf("adicionarPercentual() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_descobrirPercentual1(t * testing.T)
	{
		type args struct {
			valor       float64
			porcentagem float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := descobrirPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
					t.Errorf("descobrirPercentual() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_divisao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := divisao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("divisao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_multiplicacao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := multiplicacao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_soma(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := soma(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("soma() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_subtracao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("subtracao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_soma(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := soma(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("soma() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_soma(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := soma(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("soma() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_adicionarPercentual1(t * testing.T)
	{
		type args struct {
			valor       float64
			porcentagem float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := adicionarPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
					t.Errorf("adicionarPercentual() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_descobrirPercentual1(t * testing.T)
	{
		type args struct {
			valor       float64
			porcentagem float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := descobrirPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
					t.Errorf("descobrirPercentual() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_divisao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := divisao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("divisao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_multiplicacao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := multiplicacao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_soma(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := soma(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("soma() = %v, want %v", got, tt.want)
				}
			})
		}
	}

	func
	Test_subtracao(t * testing.T)
	{
		type args struct {
			x float64
			y float64
		}
		tests := []struct {
			name string
			args args
			want float64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
					t.Errorf("subtracao() = %v, want %v", got, tt.want)
				}
			})
		}
	}

func Test_adicionarPercentual2(t *testing.T) {
	type args struct {
		valor       float64
		porcentagem float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adicionarPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
				t.Errorf("adicionarPercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_descobrirPercentual2(t *testing.T) {
	type args struct {
		valor       float64
		porcentagem float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := descobrirPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
				t.Errorf("descobrirPercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisao1(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("divisao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplicacao1(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplicacao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_soma1(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtracao1(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("subtracao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_soma2(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_soma3(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_adicionarPercentual3(t *testing.T) {
	type args struct {
		valor       float64
		porcentagem float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := adicionarPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
				t.Errorf("adicionarPercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_descobrirPercentual3(t *testing.T) {
	type args struct {
		valor       float64
		porcentagem float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := descobrirPercentual(tt.args.valor, tt.args.porcentagem); got != tt.want {
				t.Errorf("descobrirPercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_divisao2(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("divisao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplicacao2(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplicacao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("multiplicacao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_soma4(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := soma(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("soma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtracao2(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtracao(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("subtracao() = %v, want %v", got, tt.want)
			}
		})
	}
}