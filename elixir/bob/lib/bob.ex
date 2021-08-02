defmodule Bob do
  def hey(input) do
    cond do
      question_and_shouting?(input) -> "Calm down, I know what I'm doing!"
      question?(input) -> "Sure."
      silence?(input) -> "Fine. Be that way!"
      shouting?(input) -> "Whoa, chill out!"
      true -> "Whatever."
    end
  end

  def shouting?(input) do
    input == String.upcase(input) && !numbers_only?(input)
  end

  def question?(input) do
    String.trim(input) |> String.ends_with?("?")
  end

  def silence?(input) do
    String.trim(input) == ""
  end

  def question_and_shouting?(input) do
    question?(input) && shouting?(input) && !numbers_only?(input)
  end

  def numbers_only?(input) do
    prettyify_string(input) |>
    String.trim() |>
    empty_string?
  end

  def empty_string?(input) do
    String.length(input) == 0
  end

  def prettyify_string(input) do
    String.replace(input, ~r/([0-9?:),])+/, "")
  end
end
