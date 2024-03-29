// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/rivals/v1/design_contest.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Rivals.V1 {

  /// <summary>Holder for reflection information generated from thetan/rivals/v1/design_contest.proto</summary>
  public static partial class DesignContestReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/rivals/v1/design_contest.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static DesignContestReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiV0aGV0YW4vcml2YWxzL3YxL2Rlc2lnbl9jb250ZXN0LnByb3RvEhB0aGV0",
            "YW4ucml2YWxzLnYxGiF0aGV0YW4vc2hhcmVkL3YxL2N1c3RvbWl6ZWQucHJv",
            "dG8aE3RhZ2dlci90YWdnZXIucHJvdG8itwEKBkRlc2lnbhJGCgRpdGVtGAEg",
            "ASgLMiAudGhldGFuLnNoYXJlZC52MS5JdGVtQ3VzdG9taXplZEIQmoSeAwti",
            "c29uOiJpdGVtIlIEaXRlbRIqCgZwb3NlSUQYAiABKAVCEpqEngMNYnNvbjoi",
            "cG9zZUlkIlIGcG9zZUlEEjkKC2Rlc2NyaXB0aW9uGAQgASgJQheahJ4DEmJz",
            "b246ImRlc2NyaXB0aW9uIlILZGVzY3JpcHRpb24iVgoKU3luY0Rlc2lnbhIw",
            "CgZkZXNpZ24YASABKAsyGC50aGV0YW4ucml2YWxzLnYxLkRlc2lnblIGZGVz",
            "aWduEhYKBnVzZXJJRBgCIAEoCVIGdXNlcklEQsEBChRjb20udGhldGFuLnJp",
            "dmFscy52MUISRGVzaWduQ29udGVzdFByb3RvUAFaM3RoZXRhbi1idWYvZ2Vu",
            "L2dvL3RoZXRhbi9yaXZhbHMvdjE7dGhldGFuX3JpdmFsc192MaICA1RSWKoC",
            "EFRoZXRhbi5SaXZhbHMuVjHKAhBUaGV0YW5cUml2YWxzXFYx4gIcVGhldGFu",
            "XFJpdmFsc1xWMVxHUEJNZXRhZGF0YeoCElRoZXRhbjo6Uml2YWxzOjpWMWIG",
            "cHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Thetan.Shared.V1.CustomizedReflection.Descriptor, global::Tagger.TaggerReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Rivals.V1.Design), global::Thetan.Rivals.V1.Design.Parser, new[]{ "Item", "PoseID", "Description" }, null, null, null, null),
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Rivals.V1.SyncDesign), global::Thetan.Rivals.V1.SyncDesign.Parser, new[]{ "Design", "UserID" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class Design : pb::IMessage<Design>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<Design> _parser = new pb::MessageParser<Design>(() => new Design());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<Design> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Rivals.V1.DesignContestReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Design() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Design(Design other) : this() {
      item_ = other.item_ != null ? other.item_.Clone() : null;
      poseID_ = other.poseID_;
      description_ = other.description_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public Design Clone() {
      return new Design(this);
    }

    /// <summary>Field number for the "item" field.</summary>
    public const int ItemFieldNumber = 1;
    private global::Thetan.Shared.V1.ItemCustomized item_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Shared.V1.ItemCustomized Item {
      get { return item_; }
      set {
        item_ = value;
      }
    }

    /// <summary>Field number for the "poseID" field.</summary>
    public const int PoseIDFieldNumber = 2;
    private int poseID_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int PoseID {
      get { return poseID_; }
      set {
        poseID_ = value;
      }
    }

    /// <summary>Field number for the "description" field.</summary>
    public const int DescriptionFieldNumber = 4;
    private string description_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string Description {
      get { return description_; }
      set {
        description_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as Design);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(Design other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Item, other.Item)) return false;
      if (PoseID != other.PoseID) return false;
      if (Description != other.Description) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (item_ != null) hash ^= Item.GetHashCode();
      if (PoseID != 0) hash ^= PoseID.GetHashCode();
      if (Description.Length != 0) hash ^= Description.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (item_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Item);
      }
      if (PoseID != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(PoseID);
      }
      if (Description.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(Description);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (item_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Item);
      }
      if (PoseID != 0) {
        output.WriteRawTag(16);
        output.WriteInt32(PoseID);
      }
      if (Description.Length != 0) {
        output.WriteRawTag(34);
        output.WriteString(Description);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (item_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Item);
      }
      if (PoseID != 0) {
        size += 1 + pb::CodedOutputStream.ComputeInt32Size(PoseID);
      }
      if (Description.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Description);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(Design other) {
      if (other == null) {
        return;
      }
      if (other.item_ != null) {
        if (item_ == null) {
          Item = new global::Thetan.Shared.V1.ItemCustomized();
        }
        Item.MergeFrom(other.Item);
      }
      if (other.PoseID != 0) {
        PoseID = other.PoseID;
      }
      if (other.Description.Length != 0) {
        Description = other.Description;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (item_ == null) {
              Item = new global::Thetan.Shared.V1.ItemCustomized();
            }
            input.ReadMessage(Item);
            break;
          }
          case 16: {
            PoseID = input.ReadInt32();
            break;
          }
          case 34: {
            Description = input.ReadString();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            if (item_ == null) {
              Item = new global::Thetan.Shared.V1.ItemCustomized();
            }
            input.ReadMessage(Item);
            break;
          }
          case 16: {
            PoseID = input.ReadInt32();
            break;
          }
          case 34: {
            Description = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  public sealed partial class SyncDesign : pb::IMessage<SyncDesign>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<SyncDesign> _parser = new pb::MessageParser<SyncDesign>(() => new SyncDesign());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<SyncDesign> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Rivals.V1.DesignContestReflection.Descriptor.MessageTypes[1]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SyncDesign() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SyncDesign(SyncDesign other) : this() {
      design_ = other.design_ != null ? other.design_.Clone() : null;
      userID_ = other.userID_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public SyncDesign Clone() {
      return new SyncDesign(this);
    }

    /// <summary>Field number for the "design" field.</summary>
    public const int DesignFieldNumber = 1;
    private global::Thetan.Rivals.V1.Design design_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public global::Thetan.Rivals.V1.Design Design {
      get { return design_; }
      set {
        design_ = value;
      }
    }

    /// <summary>Field number for the "userID" field.</summary>
    public const int UserIDFieldNumber = 2;
    private string userID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string UserID {
      get { return userID_; }
      set {
        userID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as SyncDesign);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(SyncDesign other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Design, other.Design)) return false;
      if (UserID != other.UserID) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (design_ != null) hash ^= Design.GetHashCode();
      if (UserID.Length != 0) hash ^= UserID.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void WriteTo(pb::CodedOutputStream output) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      output.WriteRawMessage(this);
    #else
      if (design_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Design);
      }
      if (UserID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(UserID);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalWriteTo(ref pb::WriteContext output) {
      if (design_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Design);
      }
      if (UserID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(UserID);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(ref output);
      }
    }
    #endif

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public int CalculateSize() {
      int size = 0;
      if (design_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Design);
      }
      if (UserID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(UserID);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(SyncDesign other) {
      if (other == null) {
        return;
      }
      if (other.design_ != null) {
        if (design_ == null) {
          Design = new global::Thetan.Rivals.V1.Design();
        }
        Design.MergeFrom(other.Design);
      }
      if (other.UserID.Length != 0) {
        UserID = other.UserID;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(pb::CodedInputStream input) {
    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      input.ReadRawMessage(this);
    #else
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (design_ == null) {
              Design = new global::Thetan.Rivals.V1.Design();
            }
            input.ReadMessage(Design);
            break;
          }
          case 18: {
            UserID = input.ReadString();
            break;
          }
        }
      }
    #endif
    }

    #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    void pb::IBufferMessage.InternalMergeFrom(ref pb::ParseContext input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, ref input);
            break;
          case 10: {
            if (design_ == null) {
              Design = new global::Thetan.Rivals.V1.Design();
            }
            input.ReadMessage(Design);
            break;
          }
          case 18: {
            UserID = input.ReadString();
            break;
          }
        }
      }
    }
    #endif

  }

  #endregion

}

#endregion Designer generated code
