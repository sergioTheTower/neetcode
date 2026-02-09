import pytest
from encode_decode_strings import encode, decode


@pytest.mark.parametrize(
    "input_list",
    [
        (["hello", "world"]),
        ([""]),
        (["", ""]),
        ([]),
        (["#", "##", "4#", "#abc"]),
        (["The quick brown fox", "jumps over", "the lazy dog"]),
        (["python", "🐍", "世界"]),
    ],
    ids=[
        "Standard",
        "SingleEmpty",
        "MultiEmpty",
        "EmptyList",
        "Delimiters",
        "Spaces",
        "Unicode",
    ],
)
def test_encode_decode(input_list):
    # 1. Call the standalone encode function
    encoded_str = encode(input_list)

    # 2. Call the standalone decode function
    decoded_list = decode(encoded_str)

    # 3. Assert they match
    assert decoded_list == input_list, (
        f"Failed! \n"
        f"Original: {input_list}\n"
        f"Encoded:  {encoded_str}\n"
        f"Decoded:  {decoded_list}"
    )
