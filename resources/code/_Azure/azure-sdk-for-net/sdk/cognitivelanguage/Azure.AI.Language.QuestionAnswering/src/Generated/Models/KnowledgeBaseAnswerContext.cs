// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

// <auto-generated/>

#nullable disable

namespace Azure.AI.Language.QuestionAnswering
{
    /// <summary> Context object with previous QnA's information. </summary>
    public partial class KnowledgeBaseAnswerContext
    {
        /// <summary> Initializes a new instance of <see cref="KnowledgeBaseAnswerContext"/>. </summary>
        /// <param name="previousQnaId"> Previous turn top answer result QnA ID. </param>
        public KnowledgeBaseAnswerContext(int previousQnaId)
        {
            PreviousQnaId = previousQnaId;
        }

        /// <summary> Initializes a new instance of <see cref="KnowledgeBaseAnswerContext"/>. </summary>
        /// <param name="previousQnaId"> Previous turn top answer result QnA ID. </param>
        /// <param name="previousQuestion"> Previous user query. </param>
        internal KnowledgeBaseAnswerContext(int previousQnaId, string previousQuestion)
        {
            PreviousQnaId = previousQnaId;
            PreviousQuestion = previousQuestion;
        }

        /// <summary> Previous turn top answer result QnA ID. </summary>
        public int PreviousQnaId { get; }
        /// <summary> Previous user query. </summary>
        public string PreviousQuestion { get; set; }
    }
}
