# 객체지향 설계 원칙 SOLID

- S : [SRP](#srpsingle-responsibility-principle-단일-책임-원칙) (Single Responsibility Principle, 단일 책임 원칙)
- O : [OCP](#ocpopen-closed-principle-개방-폐쇄-원칙) (Open-Closed Principle, 개방-폐쇄 원칙)
- L : LSP (Liskov Substitution Principle, 리스코프 치환 원칙)
- I : ISP (Interface Segregation Principle, 인테페이스 분리 원칙)
- D : DIP (Dependency Inversion Principle, 의존 관계 역전 원칙)

</br>

## 설계의 필요성

설계는 프로그램 코드를 이루는 각 모듈 간 의존 관계를 정의하는 것입니다.</br>
다수의 개발자들이 협업을 통해 프로그램을 만들기 때문에 좋은 설계가 중요합니다.

</br>

## 나쁜 설계

- 경직성(Rigidity) : 모듈간의 결합도(Coupling)가 너무 높아 코드를 변경하기 매우 어려운 구조.
- 부서지기 쉬움(Fragility) : 한 부분을 건드렸더니 다른 부분까지 망가지는 경우.
- 부동성(Immobility) : 코드 일부분을 현재 어플리케이션에서 분리해서 다른 프로젝트에 쓰고 싶지만 모듈 간 결합도가 너무 높아서 옮길 수 없는 경우

즉, **상호 결합도가 매우 높고 응집도가 낮은 설계**가  나쁜 설계라 할 수 있습니다.</br>
반대로 **상호 결합도가 낮고 응집도가 높은 설계**가 좋은 설계라 할 수 있겠습니다.

</br>

## SRP(Single Responsibility Principle, 단일 책임 원칙)

**모든 객체는 책임을 하나만 져야 한다** 라는 원칙 입니다.</br>
해당 원칙을 지킨 다면 코드의 재사용성을 높일 수 있습니다.

## OCP(Open-Closed Principle, 개방-폐쇄 원칙)

**확장에는 열려 있고, 변경에는 닫혀 있다** 라는 원칙 입니다.</br>
해당 원칙을 지킨 다면 상호 결합도를 줄여 새 기능을 추가할 때 기존 구현을 변경하지 않아도 됩니다.

## LSP(Liskov Substitution Principle, 리스코프 치환 원칙)

**"q(x)를 타입 T의 객체X에 대해 증명할 수 있는 속성이라 하자. 그렇다면은 S가 T의 하위 타입이라면 q(y)는 타입 S의 객체 y에 대해 증명할 수 있어야 한다."** 라는 원칙 입니다.</br>

        - T  -> S
        - q(x) => T : q(y) => S 

해당 원칙을 지키는 경우, 예상치 못한 작동을 예방할 수 있습니다.

```go
    package main

    import "fmt"

    type T interface {
        Something()
    }

    type S struct{}

    func (s *S) Something() {}  // T Interface 구현

    type U struct{}

    func (u *U) Something() {}  // T Interface 구현

    func q(t T) {
        t.Something()
    }

    var y = &S{}    // S Type의 y
    var u = &U{}    // U Type의 u

    q(y)    // S가 T의 하위 타입이기 때문에 동작해야한다.
    q(u)    // U가 T의 하위 타입이기 때문에 동작해야한다.
```

## ISP(Interface Segregation Principle, 인터페이스 분리 원칙)

**클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다.** 라는 원칙 입니다.</br>
인터페이스를 분리하면 불필요한 메서드들과 의존 관계가 끊어져 더 가볍게 인터페이스를 이용할 수 있습니다.

## DIP(Dependency Inversion Principle, 의존 관계 역전 원칙)

**상위 계층이 하위 계층에 의존하는 전통적인 의존 관계를 반전(역전)시킴으로써 상위 계층이 하위 계층의 구현으로부터 독립되게 할 수 있다.** 라는 원칙 입니다.</br>

1. 상위 모듈은 하위 모듈에 의존해서는 안 된다. 둘 다 추상 모듈에 의존해야한다.

2. 추상 모듈은 구체화된 모듈에 의존해서는 안 된다. 구체화된 모듈은 추상 모듈에 의존해야한다.

구체화된 모듈이 아닌 추상화된 모듈에 의존함으로써 확장성이 증가하며, 상호 결합도가 낮아져서 다른 프로그램으로 이식성이 증가합니다.
