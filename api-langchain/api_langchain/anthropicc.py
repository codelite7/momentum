import os

from langchain_anthropic import ChatAnthropic
from langchain_core.messages import messages_from_dict
from pyramid.response import Response

chat = ChatAnthropic(
    temperature=0,
    model="claude-3-opus-20240229"
)
system = "You are a helpful assistant."


def anthropic(request):
    try:
        messages = messages_from_dict(request.json_body)
        response = chat.invoke(messages)
        content = response.content
        return Response(content)
    except Exception as e:
        print(e)
        raise e
