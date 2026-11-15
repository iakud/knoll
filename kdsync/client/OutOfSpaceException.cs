using System.IO;

namespace Kdsync
{
    public sealed class OutOfSpaceException : IOException
    {
        internal OutOfSpaceException()
            : base("WriteContext was writing to a flat byte array and ran out of space.")
        {
        }
    }
}