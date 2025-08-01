// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * A NextDNS Rewrite into a pulumi resource
 */
export class NextDNSRewrite extends pulumi.CustomResource {
    /**
     * Get an existing NextDNSRewrite resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): NextDNSRewrite {
        return new NextDNSRewrite(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nextdns:index:NextDNSRewrite';

    /**
     * Returns true if the given object is an instance of NextDNSRewrite.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NextDNSRewrite {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NextDNSRewrite.__pulumiType;
    }

    /**
     * IP Address or Domain to rewrite domain name to.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * Domain name to apply rewrite to.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Profile Id to apply rewrite to. This overrides the default profile id.
     */
    public readonly profileId!: pulumi.Output<string | undefined>;

    /**
     * Create a NextDNSRewrite resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NextDNSRewriteArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["profileId"] = args ? args.profileId : undefined;
        } else {
            resourceInputs["content"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["profileId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NextDNSRewrite.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a NextDNSRewrite resource.
 */
export interface NextDNSRewriteArgs {
    /**
     * IP Address or Domain to rewrite domain name to.
     */
    content: pulumi.Input<string>;
    /**
     * Domain name to apply rewrite to.
     */
    name: pulumi.Input<string>;
    /**
     * Profile Id to apply rewrite to. This overrides the default profile id.
     */
    profileId?: pulumi.Input<string>;
}
