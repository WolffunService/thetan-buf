// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: thetan/multiplayer/v1/errorcode.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021, 8981
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Thetan.Multiplayer.V1 {

  /// <summary>Holder for reflection information generated from thetan/multiplayer/v1/errorcode.proto</summary>
  public static partial class ErrorcodeReflection {

    #region Descriptor
    /// <summary>File descriptor for thetan/multiplayer/v1/errorcode.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static ErrorcodeReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CiV0aGV0YW4vbXVsdGlwbGF5ZXIvdjEvZXJyb3Jjb2RlLnByb3RvEhV0aGV0",
            "YW4ubXVsdGlwbGF5ZXIudjEqKAoNQ2hhdEVycm9yQ29kZRIICgROb25lEAAS",
            "DQoJVGltZUxpbWl0EAEqkgEKDlBhcnR5RXJyb3JDb2RlEgsKB1Vua25vd24Q",
            "ABIaCg1JbnRlcm5hbEVycm9yEJ3//////////wESDwoLUGFydHlJc0Z1bGwQ",
            "ARIHCgNEbmQQAhIPCgtCbG9ja0ludml0ZRADEhAKDEZpbmRpbmdNYXRjaBAE",
            "EhoKFkRpZmZlcmVudFZlcnNpb25DbGllbnQQBUK4AQoZY29tLnRoZXRhbi5t",
            "dWx0aXBsYXllci52MUIORXJyb3Jjb2RlUHJvdG9QAVoVdGhldGFuLm11bHRp",
            "cGxheWVyLnYxogIDVE1YqgIVVGhldGFuLk11bHRpcGxheWVyLlYxygIVVGhl",
            "dGFuXE11bHRpcGxheWVyXFYx4gIhVGhldGFuXE11bHRpcGxheWVyXFYxXEdQ",
            "Qk1ldGFkYXRh6gIXVGhldGFuOjpNdWx0aXBsYXllcjo6VjFiBnByb3RvMw=="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::Thetan.Multiplayer.V1.ChatErrorCode), typeof(global::Thetan.Multiplayer.V1.PartyErrorCode), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum ChatErrorCode {
    [pbr::OriginalName("None")] None = 0,
    [pbr::OriginalName("TimeLimit")] TimeLimit = 1,
  }

  public enum PartyErrorCode {
    [pbr::OriginalName("Unknown")] Unknown = 0,
    [pbr::OriginalName("InternalError")] InternalError = -99,
    [pbr::OriginalName("PartyIsFull")] PartyIsFull = 1,
    [pbr::OriginalName("Dnd")] Dnd = 2,
    [pbr::OriginalName("BlockInvite")] BlockInvite = 3,
    /// <summary>
    /// Join party when party is finding match
    /// </summary>
    [pbr::OriginalName("FindingMatch")] FindingMatch = 4,
    /// <summary>
    ///findmatch version
    /// </summary>
    [pbr::OriginalName("DifferentVersionClient")] DifferentVersionClient = 5,
  }

  #endregion

}

#endregion Designer generated code
