﻿namespace Azure.Functions.Cli.Interfaces
{
    internal interface IProcessInfo
    {
        int Id { get; }
        string FileName { get; }
        string ProcessName { get; }
        void Kill();
    }
}
