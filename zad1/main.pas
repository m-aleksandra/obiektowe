program RandomBubbleSort;

uses
  SysUtils;

const
  MAX = 100;

type
  TIntArray = array[1..MAX] of Integer;

var
  numbers: TIntArray;
  i, count: Integer;

procedure GenerateRandomNumbers(var arr: TIntArray; count, fromVal, toVal: Integer);
var
  i: Integer;
begin
  Randomize;
  for i := 1 to count do
    arr[i] := Random(toVal - fromVal + 1) + fromVal;
end;

procedure BubbleSort(var arr: TIntArray; count: Integer);
var
  i, j, temp: Integer;
begin
  for i := 1 to count - 1 do
    for j := 1 to count - i do
      if arr[j] > arr[j + 1] then
      begin
        temp := arr[j];
        arr[j] := arr[j + 1];
        arr[j + 1] := temp;
      end;
end;

procedure DisplayNumbers(arr: TIntArray; count: Integer);
var
  i: Integer;
begin
  for i := 1 to count do
    Write(arr[i], ' ');
  Writeln;
end;

begin
  count := 50;
  GenerateRandomNumbers(numbers, count, 0, 100);

  Writeln('Numbers before sorting:');
  DisplayNumbers(numbers, count);

  BubbleSort(numbers, count);

  Writeln('Numbers after sorting:');
  DisplayNumbers(numbers, count);
end.
