// <auto-generated>
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.
// </auto-generated>

namespace Microsoft.Azure.CognitiveServices.Vision.Face
{
    using Microsoft.Rest;
    using Models;
    using System.Collections;
    using System.Collections.Generic;
    using System.IO;
    using System.Threading;
    using System.Threading.Tasks;

    /// <summary>
    /// PersonGroupPerson operations.
    /// </summary>
    public partial interface IPersonGroupPerson
    {
        /// <summary>
        /// Create a new person in a specified person group.
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='name'>
        /// User defined name, maximum length is 128.
        /// </param>
        /// <param name='userData'>
        /// User specified data. Length should not exceed 16KB.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse<Person>> CreateWithHttpMessagesAsync(string personGroupId, string name = default(string), string userData = default(string), Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// List all persons in a person group, and retrieve person information
        /// (including personId, name, userData and persistedFaceIds of
        /// registered faces of the person).
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='start'>
        /// Starting person id to return (used to list a range of persons).
        /// </param>
        /// <param name='top'>
        /// Number of persons to return starting with the person id indicated
        /// by the 'start' parameter.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse<IList<Person>>> ListWithHttpMessagesAsync(string personGroupId, string start = default(string), int? top = default(int?), Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Delete an existing person from a person group. The persistedFaceId,
        /// userData, person name and face feature in the person entry will all
        /// be deleted.
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='personId'>
        /// Id referencing a particular person.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse> DeleteWithHttpMessagesAsync(string personGroupId, System.Guid personId, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Retrieve a person's information, including registered persisted
        /// faces, name and userData.
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='personId'>
        /// Id referencing a particular person.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse<Person>> GetWithHttpMessagesAsync(string personGroupId, System.Guid personId, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Update name or userData of a person.
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='personId'>
        /// Id referencing a particular person.
        /// </param>
        /// <param name='name'>
        /// User defined name, maximum length is 128.
        /// </param>
        /// <param name='userData'>
        /// User specified data. Length should not exceed 16KB.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse> UpdateWithHttpMessagesAsync(string personGroupId, System.Guid personId, string name = default(string), string userData = default(string), Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Delete a face from a person in a person group by specified
        /// personGroupId, personId and persistedFaceId.
        /// &lt;br /&gt; Adding/deleting faces to/from a same person will be
        /// processed sequentially. Adding/deleting faces to/from different
        /// persons are processed in parallel.
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='personId'>
        /// Id referencing a particular person.
        /// </param>
        /// <param name='persistedFaceId'>
        /// Id referencing a particular persistedFaceId of an existing face.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse> DeleteFaceWithHttpMessagesAsync(string personGroupId, System.Guid personId, System.Guid persistedFaceId, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Retrieve information about a persisted face (specified by
        /// persistedFaceId, personId and its belonging personGroupId).
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='personId'>
        /// Id referencing a particular person.
        /// </param>
        /// <param name='persistedFaceId'>
        /// Id referencing a particular persistedFaceId of an existing face.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse<PersistedFace>> GetFaceWithHttpMessagesAsync(string personGroupId, System.Guid personId, System.Guid persistedFaceId, Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Add a face to a person into a person group for face identification
        /// or verification. To deal with an image contains multiple faces,
        /// input face can be specified as an image with a targetFace
        /// rectangle. It returns a persistedFaceId representing the added
        /// face. No image will be stored. Only the extracted face feature will
        /// be stored on server until [PersonGroup PersonFace -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroupperson/deleteface),
        /// [PersonGroup Person -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroupperson/delete)
        /// or [PersonGroup -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroup/delete)
        /// is called.
        /// &lt;br /&gt; Note persistedFaceId is different from faceId
        /// generated by [Face -
        /// Detect](https://docs.microsoft.com/rest/api/faceapi/face/detectwithurl).
        /// * Higher face image quality means better recognition precision.
        /// Please consider high-quality faces: frontal, clear, and face size
        /// is 200x200 pixels (100 pixels between eyes) or bigger.
        /// * Each person entry can hold up to 248 faces.
        /// * JPEG, PNG, GIF (the first frame), and BMP format are supported.
        /// The allowed image file size is from 1KB to 6MB.
        /// * "targetFace" rectangle should contain one face. Zero or multiple
        /// faces will be regarded as an error. If the provided "targetFace"
        /// rectangle is not returned from [Face -
        /// Detect](https://docs.microsoft.com/rest/api/faceapi/face/detectwithurl),
        /// there’s no guarantee to detect and add the face successfully.
        /// * Out of detectable face size (36x36 - 4096x4096 pixels), large
        /// head-pose, or large occlusions will cause failures.
        /// * Adding/deleting faces to/from a same person will be processed
        /// sequentially. Adding/deleting faces to/from different persons are
        /// processed in parallel.
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='personId'>
        /// Id referencing a particular person.
        /// </param>
        /// <param name='persistedFaceId'>
        /// Id referencing a particular persistedFaceId of an existing face.
        /// </param>
        /// <param name='userData'>
        /// User-provided data attached to the face. The size limit is 1KB.
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse> UpdateFaceWithHttpMessagesAsync(string personGroupId, System.Guid personId, System.Guid persistedFaceId, string userData = default(string), Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Add a face to a person into a person group for face identification
        /// or verification. To deal with an image contains multiple faces,
        /// input face can be specified as an image with a targetFace
        /// rectangle. It returns a persistedFaceId representing the added
        /// face. No image will be stored. Only the extracted face feature will
        /// be stored on server until [PersonGroup PersonFace -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroupperson/deleteface),
        /// [PersonGroup Person -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroupperson/delete)
        /// or [PersonGroup -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroup/delete)
        /// is called.
        /// &lt;br /&gt; Note persistedFaceId is different from faceId
        /// generated by [Face -
        /// Detect](https://docs.microsoft.com/rest/api/faceapi/face/detectwithurl).
        /// *   Higher face image quality means better recognition precision.
        /// Please consider high-quality faces: frontal, clear, and face size
        /// is 200x200 pixels (100 pixels between eyes) or bigger.
        /// *   Each person entry can hold up to 248 faces.
        /// *   JPEG, PNG, GIF (the first frame), and BMP format are supported.
        /// The allowed image file size is from 1KB to 6MB.
        /// *   "targetFace" rectangle should contain one face. Zero or
        /// multiple faces will be regarded as an error. If the provided
        /// "targetFace" rectangle is not returned from [Face -
        /// Detect](https://docs.microsoft.com/rest/api/faceapi/face/detectwithurl),
        /// there’s no guarantee to detect and add the face successfully.
        /// *   Out of detectable face size (36x36 - 4096x4096 pixels), large
        /// head-pose, or large occlusions will cause failures.
        /// *   Adding/deleting faces to/from a same person will be processed
        /// sequentially. Adding/deleting faces to/from different persons are
        /// processed in parallel.
        /// * The minimum detectable face size is 36x36 pixels in an image no
        /// larger than 1920x1080 pixels. Images with dimensions higher than
        /// 1920x1080 pixels will need a proportionally larger minimum face
        /// size.
        /// * Different 'detectionModel' values can be provided. To use and
        /// compare different detection models, please refer to [How to specify
        /// a detection
        /// model](https://docs.microsoft.com/azure/cognitive-services/face/face-api-how-to-topics/specify-detection-model).
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='personId'>
        /// Id referencing a particular person.
        /// </param>
        /// <param name='url'>
        /// Publicly reachable URL of an image
        /// </param>
        /// <param name='userData'>
        /// User-specified data about the face for any purpose. The maximum
        /// length is 1KB.
        /// </param>
        /// <param name='targetFace'>
        /// A face rectangle to specify the target face to be added to a person
        /// in the format of "targetFace=left,top,width,height". E.g.
        /// "targetFace=10,10,100,100". If there is more than one face in the
        /// image, targetFace is required to specify which face to add. No
        /// targetFace means there is only one face detected in the entire
        /// image.
        /// </param>
        /// <param name='detectionModel'>
        /// Name of detection model. Detection model is used to detect faces in
        /// the submitted image. A detection model name can be provided when
        /// performing Face - Detect or (Large)FaceList - Add Face or
        /// (Large)PersonGroup - Add Face. The default value is 'detection_01',
        /// if another model is needed, please explicitly specify it. Possible
        /// values include: 'detection_01', 'detection_02', 'detection_03'
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse<PersistedFace>> AddFaceFromUrlWithHttpMessagesAsync(string personGroupId, System.Guid personId, string url, string userData = default(string), IList<int> targetFace = default(IList<int>), string detectionModel = default(string), Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
        /// <summary>
        /// Add a face to a person into a person group for face identification
        /// or verification. To deal with an image contains multiple faces,
        /// input face can be specified as an image with a targetFace
        /// rectangle. It returns a persistedFaceId representing the added
        /// face. No image will be stored. Only the extracted face feature will
        /// be stored on server until [PersonGroup PersonFace -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroupperson/deleteface),
        /// [PersonGroup Person -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroupperson/delete)
        /// or [PersonGroup -
        /// Delete](https://docs.microsoft.com/rest/api/faceapi/persongroup/delete)
        /// is called.
        /// &lt;br /&gt; Note persistedFaceId is different from faceId
        /// generated by [Face -
        /// Detect](https://docs.microsoft.com/rest/api/faceapi/face/detectwithurl).
        /// *   Higher face image quality means better recognition precision.
        /// Please consider high-quality faces: frontal, clear, and face size
        /// is 200x200 pixels (100 pixels between eyes) or bigger.
        /// *   Each person entry can hold up to 248 faces.
        /// *   JPEG, PNG, GIF (the first frame), and BMP format are supported.
        /// The allowed image file size is from 1KB to 6MB.
        /// *   "targetFace" rectangle should contain one face. Zero or
        /// multiple faces will be regarded as an error. If the provided
        /// "targetFace" rectangle is not returned from [Face -
        /// Detect](https://docs.microsoft.com/rest/api/faceapi/face/detectwithurl),
        /// there’s no guarantee to detect and add the face successfully.
        /// *   Out of detectable face size (36x36 - 4096x4096 pixels), large
        /// head-pose, or large occlusions will cause failures.
        /// *   Adding/deleting faces to/from a same person will be processed
        /// sequentially. Adding/deleting faces to/from different persons are
        /// processed in parallel.
        /// * The minimum detectable face size is 36x36 pixels in an image no
        /// larger than 1920x1080 pixels. Images with dimensions higher than
        /// 1920x1080 pixels will need a proportionally larger minimum face
        /// size.
        /// * Different 'detectionModel' values can be provided. To use and
        /// compare different detection models, please refer to [How to specify
        /// a detection
        /// model](https://docs.microsoft.com/azure/cognitive-services/face/face-api-how-to-topics/specify-detection-model).
        /// </summary>
        /// <param name='personGroupId'>
        /// Id referencing a particular person group.
        /// </param>
        /// <param name='personId'>
        /// Id referencing a particular person.
        /// </param>
        /// <param name='image'>
        /// An image stream.
        /// </param>
        /// <param name='userData'>
        /// User-specified data about the face for any purpose. The maximum
        /// length is 1KB.
        /// </param>
        /// <param name='targetFace'>
        /// A face rectangle to specify the target face to be added to a person
        /// in the format of "targetFace=left,top,width,height". E.g.
        /// "targetFace=10,10,100,100". If there is more than one face in the
        /// image, targetFace is required to specify which face to add. No
        /// targetFace means there is only one face detected in the entire
        /// image.
        /// </param>
        /// <param name='detectionModel'>
        /// Name of detection model. Detection model is used to detect faces in
        /// the submitted image. A detection model name can be provided when
        /// performing Face - Detect or (Large)FaceList - Add Face or
        /// (Large)PersonGroup - Add Face. The default value is 'detection_01',
        /// if another model is needed, please explicitly specify it. Possible
        /// values include: 'detection_01', 'detection_02', 'detection_03'
        /// </param>
        /// <param name='customHeaders'>
        /// The headers that will be added to request.
        /// </param>
        /// <param name='cancellationToken'>
        /// The cancellation token.
        /// </param>
        /// <exception cref="APIErrorException">
        /// Thrown when the operation returned an invalid status code
        /// </exception>
        /// <exception cref="Microsoft.Rest.SerializationException">
        /// Thrown when unable to deserialize the response
        /// </exception>
        /// <exception cref="Microsoft.Rest.ValidationException">
        /// Thrown when a required parameter is null
        /// </exception>
        Task<HttpOperationResponse<PersistedFace>> AddFaceFromStreamWithHttpMessagesAsync(string personGroupId, System.Guid personId, Stream image, string userData = default(string), IList<int> targetFace = default(IList<int>), string detectionModel = default(string), Dictionary<string, List<string>> customHeaders = null, CancellationToken cancellationToken = default(CancellationToken));
    }
}
