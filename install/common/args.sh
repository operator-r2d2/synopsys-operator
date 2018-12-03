#!/bin/bash
#
# This is a rather minimal example Argbash potential
# Example taken from http://argbash.readthedocs.io/en/stable/example.html
#
# ARG_OPTIONAL_SINGLE([namespace],[n],[namespace where Synopsys operator to be installed],[synopsys-operator])
# ARG_OPTIONAL_SINGLE([docker-config],[d],[file path to Docker configuration to create the image pull secret],[])
# ARG_OPTIONAL_SINGLE([blackduck-registration-key],[k],[Black Duck registration key],[])
# ARG_OPTIONAL_SINGLE([image],[i],[Synopsys Operator image],[docker.io/blackducksoftware/synopsys-operator:2018.12.0])
# ARG_HELP([The general script's help msg])
# ARGBASH_GO()
# needed because of Argbash --> m4_ignore([
### START OF CODE GENERATED BY Argbash v2.6.1 one line above ###
# Argbash is a bash code generator used to get arguments parsing right.
# Argbash is FREE SOFTWARE, see https://argbash.io for more info
# Generated online by https://argbash.io/generate

die()
{
	local _ret=$2
	test -n "$_ret" || _ret=1
	test "$_PRINT_HELP" = yes && print_help >&2
	echo "$1" >&2
	exit ${_ret}
}

begins_with_short_option()
{
	local first_option all_short_options
	all_short_options='ndkih'
	first_option="${1:0:1}"
	test "$all_short_options" = "${all_short_options/$first_option/}" && return 1 || return 0
}

DEFAULT_FILE_PATH="../common/default-values.json"

array=( $(sed -n '/{/,/}/{s/[^:]*:[^"]*"\([^"]*\).*/\1/p;}' "$DEFAULT_FILE_PATH") ) 
NS="${array[0]}"
IMAGE="${array[1]}"
REG_KEY="${array[2]}"
DOCKER_CONFIG_PATH="${array[3]}"

# THE DEFAULTS INITIALIZATION - OPTIONALS
_arg_namespace="$NS"
_arg_docker_config="$DOCKER_CONFIG_PATH"
_arg_blackduck_registration_key="$REG_KEY"
_arg_image="$IMAGE"

print_help ()
{
	printf '%s\n' "The general script's help msg"
	printf 'Usage: %s [-n|--namespace <arg>] [-d|--docker-config <arg>] [-k|--blackduck-registration-key <arg>] [-i|--image <arg>] [-h|--help]\n' "$0"
	printf '\t%s\n' "-n,--namespace: namespace where Synopsys operator to be installed (default: '$NS')"
	printf '\t%s\n' "-i,--image: Synopsys Operator image (default: '$IMAGE')"
        printf '\t%s\n' "-k,--blackduck-registration-key: Black Duck registration key (default: '$REG_KEY')"
        printf '\t%s\n' "-d,--docker-config: file path to Docker configuration to create the image pull secret (default: '$DOCKER_CONFIG_PATH')"
	printf '\t%s\n' "-h,--help: Prints help"
}

parse_commandline ()
{
	while test $# -gt 0
	do
		_key="$1"
		case "$_key" in
			-n|--namespace)
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
				_arg_namespace="$2"
				shift
				;;
			--namespace=*)
				_arg_namespace="${_key##--namespace=}"
				;;
			-n*)
				_arg_namespace="${_key##-n}"
				;;
			-d|--docker-config)
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
				_arg_docker_config="$2"
				shift
				;;
			--docker-config=*)
				_arg_docker_config="${_key##--docker-config=}"
				;;
			-d*)
				_arg_docker_config="${_key##-d}"
				;;
			-k|--blackduck-registration-key)
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
				_arg_blackduck_registration_key="$2"
				shift
				;;
			--blackduck-registration-key=*)
				_arg_blackduck_registration_key="${_key##--blackduck-registration-key=}"
				;;
			-k*)
				_arg_blackduck_registration_key="${_key##-k}"
				;;
			-i|--image)
				test $# -lt 2 && die "Missing value for the optional argument '$_key'." 1
				_arg_image="$2"
				shift
				;;
			--image=*)
				_arg_image="${_key##--image=}"
				;;
			-i*)
				_arg_image="${_key##-i}"
				;;
			-h|--help)
				print_help
				exit 0
				;;
			-h*)
				print_help
				exit 0
				;;
			*)
				_PRINT_HELP=yes die "FATAL ERROR: Got an unexpected argument '$1'" 1
				;;
		esac
		shift
	done
}

parse_commandline "$@"

# OTHER STUFF GENERATED BY Argbash

### END OF CODE GENERATED BY Argbash (sortof) ### ])
