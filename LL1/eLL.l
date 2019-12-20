E->T E1
E1->op1 T E1
E1->$
T->F T1
T1->op2 F T1
T1->$
F->( E )
F->-F
F->I
I->var
I->const
op1->+
op1->-
op2->*
op2->/