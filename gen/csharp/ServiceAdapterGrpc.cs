// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/world/v1/service_adapter.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Thetan.World.V1 {
  public static partial class ThetanWorldAdapterService
  {
    static readonly string __ServiceName = "thetan.world.v1.ThetanWorldAdapterService";

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static void __Helper_SerializeMessage(global::Google.Protobuf.IMessage message, grpc::SerializationContext context)
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (message is global::Google.Protobuf.IBufferMessage)
      {
        context.SetPayloadLength(message.CalculateSize());
        global::Google.Protobuf.MessageExtensions.WriteTo(message, context.GetBufferWriter());
        context.Complete();
        return;
      }
      #endif
      context.Complete(global::Google.Protobuf.MessageExtensions.ToByteArray(message));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static class __Helper_MessageCache<T>
    {
      public static readonly bool IsBufferMessage = global::System.Reflection.IntrospectionExtensions.GetTypeInfo(typeof(global::Google.Protobuf.IBufferMessage)).IsAssignableFrom(typeof(T));
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static T __Helper_DeserializeMessage<T>(grpc::DeserializationContext context, global::Google.Protobuf.MessageParser<T> parser) where T : global::Google.Protobuf.IMessage<T>
    {
      #if !GRPC_DISABLE_PROTOBUF_BUFFER_SERIALIZATION
      if (__Helper_MessageCache<T>.IsBufferMessage)
      {
        return parser.ParseFrom(context.PayloadAsReadOnlySequence());
      }
      #endif
      return parser.ParseFrom(context.PayloadAsNewBuffer());
    }

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Google.Protobuf.WellKnownTypes.Empty> __Marshaller_google_protobuf_Empty = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Google.Protobuf.WellKnownTypes.Empty.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.AvailableItem> __Marshaller_thetan_world_v1_AvailableItem = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.AvailableItem.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.IsValidItemsRequest> __Marshaller_thetan_world_v1_IsValidItemsRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.IsValidItemsRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.CheckItemResponse> __Marshaller_thetan_world_v1_CheckItemResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.CheckItemResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.SendItemsRequest> __Marshaller_thetan_world_v1_SendItemsRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.SendItemsRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.SendItemsResponse> __Marshaller_thetan_world_v1_SendItemsResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.SendItemsResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.CreateItemRequest> __Marshaller_thetan_world_v1_CreateItemRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.CreateItemRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.ItemResponse> __Marshaller_thetan_world_v1_ItemResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.ItemResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.GetItemsRequest> __Marshaller_thetan_world_v1_GetItemsRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.GetItemsRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.GetItemsResponse> __Marshaller_thetan_world_v1_GetItemsResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.GetItemsResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.InUsedRequest> __Marshaller_thetan_world_v1_InUsedRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.InUsedRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.World.V1.InUsedResponse> __Marshaller_thetan_world_v1_InUsedResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.World.V1.InUsedResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Google.Protobuf.WellKnownTypes.Empty, global::Thetan.World.V1.AvailableItem> __Method_GetAvailableItems = new grpc::Method<global::Google.Protobuf.WellKnownTypes.Empty, global::Thetan.World.V1.AvailableItem>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetAvailableItems",
        __Marshaller_google_protobuf_Empty,
        __Marshaller_thetan_world_v1_AvailableItem);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.World.V1.IsValidItemsRequest, global::Thetan.World.V1.CheckItemResponse> __Method_IsValidItems = new grpc::Method<global::Thetan.World.V1.IsValidItemsRequest, global::Thetan.World.V1.CheckItemResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "IsValidItems",
        __Marshaller_thetan_world_v1_IsValidItemsRequest,
        __Marshaller_thetan_world_v1_CheckItemResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.World.V1.SendItemsRequest, global::Thetan.World.V1.SendItemsResponse> __Method_SendItems = new grpc::Method<global::Thetan.World.V1.SendItemsRequest, global::Thetan.World.V1.SendItemsResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "SendItems",
        __Marshaller_thetan_world_v1_SendItemsRequest,
        __Marshaller_thetan_world_v1_SendItemsResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.World.V1.CreateItemRequest, global::Thetan.World.V1.ItemResponse> __Method_CreateNFTItem = new grpc::Method<global::Thetan.World.V1.CreateItemRequest, global::Thetan.World.V1.ItemResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateNFTItem",
        __Marshaller_thetan_world_v1_CreateItemRequest,
        __Marshaller_thetan_world_v1_ItemResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.World.V1.GetItemsRequest, global::Thetan.World.V1.GetItemsResponse> __Method_GetItems = new grpc::Method<global::Thetan.World.V1.GetItemsRequest, global::Thetan.World.V1.GetItemsResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "GetItems",
        __Marshaller_thetan_world_v1_GetItemsRequest,
        __Marshaller_thetan_world_v1_GetItemsResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.World.V1.InUsedRequest, global::Thetan.World.V1.InUsedResponse> __Method_IsInUsed = new grpc::Method<global::Thetan.World.V1.InUsedRequest, global::Thetan.World.V1.InUsedResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "IsInUsed",
        __Marshaller_thetan_world_v1_InUsedRequest,
        __Marshaller_thetan_world_v1_InUsedResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.World.V1.ServiceAdapterReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of ThetanWorldAdapterService</summary>
    [grpc::BindServiceMethod(typeof(ThetanWorldAdapterService), "BindService")]
    public abstract partial class ThetanWorldAdapterServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.World.V1.AvailableItem> GetAvailableItems(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.World.V1.CheckItemResponse> IsValidItems(global::Thetan.World.V1.IsValidItemsRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.World.V1.SendItemsResponse> SendItems(global::Thetan.World.V1.SendItemsRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.World.V1.ItemResponse> CreateNFTItem(global::Thetan.World.V1.CreateItemRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.World.V1.GetItemsResponse> GetItems(global::Thetan.World.V1.GetItemsRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.World.V1.InUsedResponse> IsInUsed(global::Thetan.World.V1.InUsedRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for ThetanWorldAdapterService</summary>
    public partial class ThetanWorldAdapterServiceClient : grpc::ClientBase<ThetanWorldAdapterServiceClient>
    {
      /// <summary>Creates a new client for ThetanWorldAdapterService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public ThetanWorldAdapterServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for ThetanWorldAdapterService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public ThetanWorldAdapterServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected ThetanWorldAdapterServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected ThetanWorldAdapterServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.AvailableItem GetAvailableItems(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetAvailableItems(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.AvailableItem GetAvailableItems(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetAvailableItems, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.AvailableItem> GetAvailableItemsAsync(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetAvailableItemsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.AvailableItem> GetAvailableItemsAsync(global::Google.Protobuf.WellKnownTypes.Empty request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetAvailableItems, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.CheckItemResponse IsValidItems(global::Thetan.World.V1.IsValidItemsRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return IsValidItems(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.CheckItemResponse IsValidItems(global::Thetan.World.V1.IsValidItemsRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_IsValidItems, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.CheckItemResponse> IsValidItemsAsync(global::Thetan.World.V1.IsValidItemsRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return IsValidItemsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.CheckItemResponse> IsValidItemsAsync(global::Thetan.World.V1.IsValidItemsRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_IsValidItems, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.SendItemsResponse SendItems(global::Thetan.World.V1.SendItemsRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SendItems(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.SendItemsResponse SendItems(global::Thetan.World.V1.SendItemsRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_SendItems, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.SendItemsResponse> SendItemsAsync(global::Thetan.World.V1.SendItemsRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return SendItemsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.SendItemsResponse> SendItemsAsync(global::Thetan.World.V1.SendItemsRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_SendItems, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.ItemResponse CreateNFTItem(global::Thetan.World.V1.CreateItemRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateNFTItem(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.ItemResponse CreateNFTItem(global::Thetan.World.V1.CreateItemRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateNFTItem, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.ItemResponse> CreateNFTItemAsync(global::Thetan.World.V1.CreateItemRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateNFTItemAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.ItemResponse> CreateNFTItemAsync(global::Thetan.World.V1.CreateItemRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateNFTItem, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.GetItemsResponse GetItems(global::Thetan.World.V1.GetItemsRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetItems(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.GetItemsResponse GetItems(global::Thetan.World.V1.GetItemsRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_GetItems, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.GetItemsResponse> GetItemsAsync(global::Thetan.World.V1.GetItemsRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return GetItemsAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.GetItemsResponse> GetItemsAsync(global::Thetan.World.V1.GetItemsRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_GetItems, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.InUsedResponse IsInUsed(global::Thetan.World.V1.InUsedRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return IsInUsed(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.World.V1.InUsedResponse IsInUsed(global::Thetan.World.V1.InUsedRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_IsInUsed, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.InUsedResponse> IsInUsedAsync(global::Thetan.World.V1.InUsedRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return IsInUsedAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.World.V1.InUsedResponse> IsInUsedAsync(global::Thetan.World.V1.InUsedRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_IsInUsed, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override ThetanWorldAdapterServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new ThetanWorldAdapterServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(ThetanWorldAdapterServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_GetAvailableItems, serviceImpl.GetAvailableItems)
          .AddMethod(__Method_IsValidItems, serviceImpl.IsValidItems)
          .AddMethod(__Method_SendItems, serviceImpl.SendItems)
          .AddMethod(__Method_CreateNFTItem, serviceImpl.CreateNFTItem)
          .AddMethod(__Method_GetItems, serviceImpl.GetItems)
          .AddMethod(__Method_IsInUsed, serviceImpl.IsInUsed).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, ThetanWorldAdapterServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_GetAvailableItems, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Google.Protobuf.WellKnownTypes.Empty, global::Thetan.World.V1.AvailableItem>(serviceImpl.GetAvailableItems));
      serviceBinder.AddMethod(__Method_IsValidItems, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.World.V1.IsValidItemsRequest, global::Thetan.World.V1.CheckItemResponse>(serviceImpl.IsValidItems));
      serviceBinder.AddMethod(__Method_SendItems, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.World.V1.SendItemsRequest, global::Thetan.World.V1.SendItemsResponse>(serviceImpl.SendItems));
      serviceBinder.AddMethod(__Method_CreateNFTItem, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.World.V1.CreateItemRequest, global::Thetan.World.V1.ItemResponse>(serviceImpl.CreateNFTItem));
      serviceBinder.AddMethod(__Method_GetItems, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.World.V1.GetItemsRequest, global::Thetan.World.V1.GetItemsResponse>(serviceImpl.GetItems));
      serviceBinder.AddMethod(__Method_IsInUsed, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.World.V1.InUsedRequest, global::Thetan.World.V1.InUsedResponse>(serviceImpl.IsInUsed));
    }

  }
}
#endregion
