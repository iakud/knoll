namespace Kdsync
{
    internal struct WriterInternalState
    {
        internal int limit;  // the size of the current buffer
        internal int position;  // position in the current buffer

        internal WriteBufferHelper writeBufferHelper;
    }
}