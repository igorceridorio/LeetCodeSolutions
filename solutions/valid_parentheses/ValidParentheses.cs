public class Solution {
    public bool IsValid(string s) {
        Stack<char> stack = new Stack<char>();

            while(s.Length != 0)
            {
                char c = Convert.ToChar(s[0]);
                s = s.Substring(1);

                if (c == '(' || c == '[' || c == '{')
                {
                    stack.Push(c);
                }
                else if ((c == ')' || c == ']' || c == '}') && stack.Count() == 0)
                {
                    return false;
                }
                else if (c == ')' && stack.Count() > 0)
                {
                    if (stack.Peek() == '(')
                    {
                        stack.Pop();
                    }
                    else
                    {
                        return false;
                    }

                }
                else if (c == ']' && stack.Count() > 0)
                {
                    if (stack.Peek() == '[')
                    {
                        stack.Pop();
                    }
                    else
                    {
                        return false;
                    }

                }
                else if (c == '}' && stack.Count() > 0)
                {
                    if (stack.Peek() == '{')
                    {
                        stack.Pop();
                    }
                    else
                    {
                        return false;
                    }

                }

            }

            if (stack.Any())
            {
                return false;
            }

            return true;
    }
}