# frozen_string_literal: true

# Q. この関数は何をしているのか？のように関数の説明をCopilotに聞いてみましょう。
def fizzbuzz(num)
  if (num % 15).zero?
    'FizzBuzz'
  elsif (num % 5).zero?
    'Buzz'
  elsif (num % 3).zero?
    'Fizz'
  else
    num
  end
end

puts 'Please input number'
puts fizzbuzz(gets.chomp.to_i)
