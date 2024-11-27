// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/immortal/v1/immortal_match.proto
// </auto-generated>
#pragma warning disable 0414, 1591, 8981, 0612
#region Designer generated code

using grpc = global::Grpc.Core;

namespace Thetan.Immortal.V1 {
  public static partial class MatchDirectorService
  {
    static readonly string __ServiceName = "thetan.immortal.v1.MatchDirectorService";

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
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.CancelTicketRequest> __Marshaller_thetan_immortal_v1_CancelTicketRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.CancelTicketRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.CancelTicketResponse> __Marshaller_thetan_immortal_v1_CancelTicketResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.CancelTicketResponse.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest> __Marshaller_thetan_immortal_v1_CreateMatchNonMatchingRequest = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest.Parser));
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Marshaller<global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse> __Marshaller_thetan_immortal_v1_CreateMatchNonMatchingResponse = grpc::Marshallers.Create(__Helper_SerializeMessage, context => __Helper_DeserializeMessage(context, global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse.Parser));

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Immortal.V1.CancelTicketRequest, global::Thetan.Immortal.V1.CancelTicketResponse> __Method_CancelTicket = new grpc::Method<global::Thetan.Immortal.V1.CancelTicketRequest, global::Thetan.Immortal.V1.CancelTicketResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CancelTicket",
        __Marshaller_thetan_immortal_v1_CancelTicketRequest,
        __Marshaller_thetan_immortal_v1_CancelTicketResponse);

    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    static readonly grpc::Method<global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest, global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse> __Method_CreateMatchNonMatching = new grpc::Method<global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest, global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse>(
        grpc::MethodType.Unary,
        __ServiceName,
        "CreateMatchNonMatching",
        __Marshaller_thetan_immortal_v1_CreateMatchNonMatchingRequest,
        __Marshaller_thetan_immortal_v1_CreateMatchNonMatchingResponse);

    /// <summary>Service descriptor</summary>
    public static global::Google.Protobuf.Reflection.ServiceDescriptor Descriptor
    {
      get { return global::Thetan.Immortal.V1.ImmortalMatchReflection.Descriptor.Services[0]; }
    }

    /// <summary>Base class for server-side implementations of MatchDirectorService</summary>
    [grpc::BindServiceMethod(typeof(MatchDirectorService), "BindService")]
    public abstract partial class MatchDirectorServiceBase
    {
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Immortal.V1.CancelTicketResponse> CancelTicket(global::Thetan.Immortal.V1.CancelTicketRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::System.Threading.Tasks.Task<global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse> CreateMatchNonMatching(global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest request, grpc::ServerCallContext context)
      {
        throw new grpc::RpcException(new grpc::Status(grpc::StatusCode.Unimplemented, ""));
      }

    }

    /// <summary>Client for MatchDirectorService</summary>
    public partial class MatchDirectorServiceClient : grpc::ClientBase<MatchDirectorServiceClient>
    {
      /// <summary>Creates a new client for MatchDirectorService</summary>
      /// <param name="channel">The channel to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public MatchDirectorServiceClient(grpc::ChannelBase channel) : base(channel)
      {
      }
      /// <summary>Creates a new client for MatchDirectorService that uses a custom <c>CallInvoker</c>.</summary>
      /// <param name="callInvoker">The callInvoker to use to make remote calls.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public MatchDirectorServiceClient(grpc::CallInvoker callInvoker) : base(callInvoker)
      {
      }
      /// <summary>Protected parameterless constructor to allow creation of test doubles.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected MatchDirectorServiceClient() : base()
      {
      }
      /// <summary>Protected constructor to allow creation of configured clients.</summary>
      /// <param name="configuration">The client configuration.</param>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected MatchDirectorServiceClient(ClientBaseConfiguration configuration) : base(configuration)
      {
      }

      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.CancelTicketResponse CancelTicket(global::Thetan.Immortal.V1.CancelTicketRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CancelTicket(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.CancelTicketResponse CancelTicket(global::Thetan.Immortal.V1.CancelTicketRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CancelTicket, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.CancelTicketResponse> CancelTicketAsync(global::Thetan.Immortal.V1.CancelTicketRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CancelTicketAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.CancelTicketResponse> CancelTicketAsync(global::Thetan.Immortal.V1.CancelTicketRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CancelTicket, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse CreateMatchNonMatching(global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateMatchNonMatching(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse CreateMatchNonMatching(global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest request, grpc::CallOptions options)
      {
        return CallInvoker.BlockingUnaryCall(__Method_CreateMatchNonMatching, null, options, request);
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse> CreateMatchNonMatchingAsync(global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest request, grpc::Metadata headers = null, global::System.DateTime? deadline = null, global::System.Threading.CancellationToken cancellationToken = default(global::System.Threading.CancellationToken))
      {
        return CreateMatchNonMatchingAsync(request, new grpc::CallOptions(headers, deadline, cancellationToken));
      }
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      public virtual grpc::AsyncUnaryCall<global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse> CreateMatchNonMatchingAsync(global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest request, grpc::CallOptions options)
      {
        return CallInvoker.AsyncUnaryCall(__Method_CreateMatchNonMatching, null, options, request);
      }
      /// <summary>Creates a new instance of client from given <c>ClientBaseConfiguration</c>.</summary>
      [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
      protected override MatchDirectorServiceClient NewInstance(ClientBaseConfiguration configuration)
      {
        return new MatchDirectorServiceClient(configuration);
      }
    }

    /// <summary>Creates service definition that can be registered with a server</summary>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static grpc::ServerServiceDefinition BindService(MatchDirectorServiceBase serviceImpl)
    {
      return grpc::ServerServiceDefinition.CreateBuilder()
          .AddMethod(__Method_CancelTicket, serviceImpl.CancelTicket)
          .AddMethod(__Method_CreateMatchNonMatching, serviceImpl.CreateMatchNonMatching).Build();
    }

    /// <summary>Register service method with a service binder with or without implementation. Useful when customizing the service binding logic.
    /// Note: this method is part of an experimental API that can change or be removed without any prior notice.</summary>
    /// <param name="serviceBinder">Service methods will be bound by calling <c>AddMethod</c> on this object.</param>
    /// <param name="serviceImpl">An object implementing the server-side handling logic.</param>
    [global::System.CodeDom.Compiler.GeneratedCode("grpc_csharp_plugin", null)]
    public static void BindService(grpc::ServiceBinderBase serviceBinder, MatchDirectorServiceBase serviceImpl)
    {
      serviceBinder.AddMethod(__Method_CancelTicket, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Immortal.V1.CancelTicketRequest, global::Thetan.Immortal.V1.CancelTicketResponse>(serviceImpl.CancelTicket));
      serviceBinder.AddMethod(__Method_CreateMatchNonMatching, serviceImpl == null ? null : new grpc::UnaryServerMethod<global::Thetan.Immortal.V1.CreateMatchNonMatchingRequest, global::Thetan.Immortal.V1.CreateMatchNonMatchingResponse>(serviceImpl.CreateMatchNonMatching));
    }

  }
}
#endregion
