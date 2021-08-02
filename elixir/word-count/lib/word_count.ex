defmodule WordCount do
  @regex Regex.compile("/[^a-zA-Z0-9-äöüÄÖÜß]/")
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    {status, regex} = @regex

    sentence |>
      String.downcase() |>
      String.replace(regex, " ") |>
      String.split() |>
      Enum.reduce(Map.new(), fn x, acc -> Map.update(acc, x, 1, &(&1 + 1)) end)
  end
end
