// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/rivals/v1/add_in.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Rivals.V1 {

  /// <summary>Holder for reflection information generated from thetan/rivals/v1/add_in.proto</summary>
  public static partial class AddInReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/rivals/v1/add_in.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static AddInReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Ch10aGV0YW4vcml2YWxzL3YxL2FkZF9pbi5wcm90bxIQdGhldGFuLnJpdmFs",
            "cy52MSJnCgtNaWdyYXRlSXRlbRIcCglvbGRJdGVtSUQYASABKAlSCW9sZEl0",
            "ZW1JRBIcCgluZXdJdGVtSUQYAiABKAlSCW5ld0l0ZW1JRBIcCgluZXdVc2Vy",
            "SUQYAyABKAlSCW5ld1VzZXJJREK5AQoUY29tLnRoZXRhbi5yaXZhbHMudjFC",
            "CkFkZEluUHJvdG9QAVozdGhldGFuLWJ1Zi9nZW4vZ28vdGhldGFuL3JpdmFs",
            "cy92MTt0aGV0YW5fcml2YWxzX3YxogIDVFJYqgIQVGhldGFuLlJpdmFscy5W",
            "McoCEFRoZXRhblxSaXZhbHNcVjHiAhxUaGV0YW5cUml2YWxzXFYxXEdQQk1l",
            "dGFkYXRh6gISVGhldGFuOjpSaXZhbHM6OlYxYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Thetan.Rivals.V1.MigrateItem), global::Thetan.Rivals.V1.MigrateItem.Parser, new[]{ "OldItemID", "NewItemID", "NewUserID" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  /// <summary>
  /// Deprecated
  /// </summary>
  public sealed partial class MigrateItem : pb::IMessage<MigrateItem>
  #if !GOOGLE_PROTOBUF_REFSTRUCT_COMPATIBILITY_MODE
      , pb::IBufferMessage
  #endif
  {
    private static readonly pb::MessageParser<MigrateItem> _parser = new pb::MessageParser<MigrateItem>(() => new MigrateItem());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pb::MessageParser<MigrateItem> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Thetan.Rivals.V1.AddInReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MigrateItem() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MigrateItem(MigrateItem other) : this() {
      oldItemID_ = other.oldItemID_;
      newItemID_ = other.newItemID_;
      newUserID_ = other.newUserID_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public MigrateItem Clone() {
      return new MigrateItem(this);
    }

    /// <summary>Field number for the "oldItemID" field.</summary>
    public const int OldItemIDFieldNumber = 1;
    private string oldItemID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string OldItemID {
      get { return oldItemID_; }
      set {
        oldItemID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "newItemID" field.</summary>
    public const int NewItemIDFieldNumber = 2;
    private string newItemID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string NewItemID {
      get { return newItemID_; }
      set {
        newItemID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "newUserID" field.</summary>
    public const int NewUserIDFieldNumber = 3;
    private string newUserID_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public string NewUserID {
      get { return newUserID_; }
      set {
        newUserID_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override bool Equals(object other) {
      return Equals(other as MigrateItem);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public bool Equals(MigrateItem other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (OldItemID != other.OldItemID) return false;
      if (NewItemID != other.NewItemID) return false;
      if (NewUserID != other.NewUserID) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public override int GetHashCode() {
      int hash = 1;
      if (OldItemID.Length != 0) hash ^= OldItemID.GetHashCode();
      if (NewItemID.Length != 0) hash ^= NewItemID.GetHashCode();
      if (NewUserID.Length != 0) hash ^= NewUserID.GetHashCode();
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
      if (OldItemID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(OldItemID);
      }
      if (NewItemID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(NewItemID);
      }
      if (NewUserID.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(NewUserID);
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
      if (OldItemID.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(OldItemID);
      }
      if (NewItemID.Length != 0) {
        output.WriteRawTag(18);
        output.WriteString(NewItemID);
      }
      if (NewUserID.Length != 0) {
        output.WriteRawTag(26);
        output.WriteString(NewUserID);
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
      if (OldItemID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(OldItemID);
      }
      if (NewItemID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(NewItemID);
      }
      if (NewUserID.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(NewUserID);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    [global::System.CodeDom.Compiler.GeneratedCode("protoc", null)]
    public void MergeFrom(MigrateItem other) {
      if (other == null) {
        return;
      }
      if (other.OldItemID.Length != 0) {
        OldItemID = other.OldItemID;
      }
      if (other.NewItemID.Length != 0) {
        NewItemID = other.NewItemID;
      }
      if (other.NewUserID.Length != 0) {
        NewUserID = other.NewUserID;
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
            OldItemID = input.ReadString();
            break;
          }
          case 18: {
            NewItemID = input.ReadString();
            break;
          }
          case 26: {
            NewUserID = input.ReadString();
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
            OldItemID = input.ReadString();
            break;
          }
          case 18: {
            NewItemID = input.ReadString();
            break;
          }
          case 26: {
            NewUserID = input.ReadString();
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
